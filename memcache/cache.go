package memcache

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	ErrCacheMiss    = errors.New("memcache: cache miss")                                      // ErrCacheMiss means that a Get failed because the item wasn't present.
	ErrCASConflict  = errors.New("memcache: compare-and-swap conflict")                       // ErrCASConflict means that a CAS operation failed due to the provided comparison value not matching the existing value.
	ErrNotStored    = errors.New("memcache: item not stored")                                 // ErrNotStored means that a Set or Add operation was not stored.
	ErrServerError  = errors.New("memcache: server error")                                    // ErrServerError means that a server error occurred.
	ErrNoStats      = errors.New("memcache: no statistics available")                         // ErrNoStats means that no statistics are available.
	ErrMalformedKey = errors.New("malformed: key is too long or contains invalid characters") // ErrMalformedKey means that the key is too long or contains invalid characters.
	ErrNoServers    = errors.New("memcache: no servers configured or available")              // ErrNoServers means that no servers are available to resolve the item's key.
)

const (
	DefaultTimeout      = 500 * time.Millisecond
	DefaultMaxIdleConns = 2
)

const buffered = 8

func resumableError(err error) bool {
	switch err {
	case ErrCacheMiss, ErrCASConflict, ErrNotStored, ErrMalformedKey:
		return true
	}
	return false
}

func legalKey(key string) bool {
	if len(key) > 250 {
		return false
	}
	for i := 0; i < len(key); i++ {
		if key[i] <= ' ' || key[i] == 0x7f {
			return false
		}
	}
	return true
}

var (
	crlf            = []byte("\r\n")
	space           = []byte(" ")
	resultOK        = []byte("OK\r\n")
	resultStored    = []byte("STORED\r\n")
	resultNotStored = []byte("NOT_STORED\r\n")
	resultExists    = []byte("EXISTS\r\n")
	resultNotFound  = []byte("NOT_FOUND\r\n")
	resultDeleted   = []byte("DELETED\r\n")
	resultEnd       = []byte("END\r\n")
	resultOk        = []byte("OK\r\n")
	resultTouched   = []byte("TOUCHED\r\n")

	resultClientErrorPrefix = []byte("CLIENT_ERROR ")
	versionPrefix           = []byte("VERSION")
)

func New(server ...string) *Client {
	ss := new(ServerList)
	ss.SetServers(server...)
	return NewFromSelector(ss)
}

func NewFromSelector(ss ServerSelector) *Client {
	return &Client{selector: ss}
}

// Memcache consumer for connection of the memcache pool.
type Client struct {
	DialContext func(ctx context.Context, network, address string) (net.Conn, error)

	Timeout time.Duration

	MaxIdleConns int

	selector ServerSelector

	mutx     sync.Mutex
	freeconn map[string][]*conn
}

// Memcache store item in the memcache pool.
type Item struct {
	// Unique key of this item in the memcache pool. (limited to 250 bytes)
	Key string

	// Value of Item
	Value []byte

	// Flags are server-opaque 32-bit unsigned values set by the client.
	Flags uint32

	// Expiration time in seconds. (0 means never expire)
	Expiration int32

	// CasID is a unique 64-bit value set by the server when the item is stored.
	CasID uint64
}

type conn struct {
	nc   net.Conn
	rw   *bufio.ReadWriter
	addr net.Addr
	c    *Client
}

func (cn *conn) release() {
	cn.c.putFreeConn(cn.addr, cn)
}

func (cn *conn) extendDeadline() {
	cn.nc.SetDeadline(time.Now().Add(cn.c.netTimeout()))
}

func (cn *conn) condRelease(err *error) {
	if *err == nil || resumableError(*err) {
		cn.release()
	} else {
		cn.nc.Close()
	}
}

func (c *Client) putFreeConn(addr net.Addr, cn *conn) {

	c.mutx.Lock()
	defer c.mutx.Unlock()

	if c.freeconn == nil {
		c.freeconn = make(map[string][]*conn)
	}

	freelist := c.freeconn[addr.String()]

	if len(freelist) >= c.maxIdleConns() {
		cn.nc.Close()
		return
	}

	c.freeconn[addr.String()] = append(freelist, cn)
}

func (c *Client) getFreeConn(addr net.Addr) (cn *conn, ok bool) {
	c.mutx.Lock()
	defer c.mutx.Unlock()

	if c.freeconn == nil {
		return nil, false
	}

	freelist, ok := c.freeconn[addr.String()]

	if !ok || len(freelist) == 0 {
		return nil, false
	}

	cn = freelist[len(freelist)-1]

	c.freeconn[addr.String()] = freelist[:len(freelist)-1]

	return cn, true
}

func (c *Client) netTimeout() time.Duration {

	if c.Timeout != 0 {
		return c.Timeout
	}

	return DefaultTimeout
}

func (c *Client) maxIdleConns() int {

	if c.MaxIdleConns > 0 {
		return c.MaxIdleConns
	}

	return DefaultMaxIdleConns
}

type ConnectTimeoutError struct {
	Addr net.Addr
}

func (cte *ConnectTimeoutError) Error() string {
	return "memcache: connect timeout to " + cte.Addr.String()
}

func (c *Client) dial(addr net.Addr) (net.Conn, error) {

	ctx, cancel := context.WithTimeout(context.Background(), c.netTimeout())

	defer cancel()

	dialerContext := c.DialContext
	if dialerContext == nil {
		dialer := net.Dialer{
			Timeout: c.netTimeout(),
		}
		dialerContext = dialer.DialContext
	}

	nc, err := dialerContext(ctx, addr.Network(), addr.String())

	if err == nil {
		return nc, nil
	}

	if ne, ok := err.(net.Error); ok && ne.Timeout() {
		return nil, &ConnectTimeoutError{addr}
	}

	return nil, err
}

func (c *Client) getConn(addr net.Addr) (*conn, error) {

	cn, ok := c.getFreeConn(addr)

	if ok {
		cn.extendDeadline()
		return cn, nil
	}

	nc, err := c.dial(addr)

	if err != nil {
		return nil, err
	}

	cn = &conn{
		nc:   nc,
		addr: addr,
		rw:   bufio.NewReadWriter(bufio.NewReader(nc), bufio.NewWriter(nc)),
		c:    c,
	}

	cn.extendDeadline()

	return cn, nil
}

func (c *Client) onItem(item *Item, fn func(*Client, *bufio.ReadWriter, *Item) error) error {

	addr, err := c.selector.PickServer(item.Key)

	if err != nil {
		return err
	}

	cn, err := c.getConn(addr)

	if err != nil {
		return err
	}

	defer cn.condRelease(&err)

	if err = fn(c, cn.rw, item); err != nil {
		return err
	}

	return nil
}

func (c *Client) FlushAll() error {
	return c.selector.Each(c.flushAllFromAddr)
}

func (c *Client) Get(key string) (item *Item, err error) {

	err = c.withKeyAddr(key, func(addr net.Addr) error {
		return c.getFromAddr(addr, []string{key}, func(it *Item) { item = it })
	})

	if err == nil && item == nil {
		err = ErrCacheMiss
	}

	return
}

func (c *Client) Touch(key string, seconds int32) (err error) {
	return c.withKeyAddr(key, func(addr net.Addr) error {
		return c.touchFromAddr(addr, []string{key}, seconds)
	})
}

func (c *Client) withKeyAddr(key string, fn func(net.Addr) error) (err error) {

	if !legalKey(key) {
		return ErrMalformedKey
	}

	addr, err := c.selector.PickServer(key)

	if err != nil {
		return err
	}

	return fn(addr)
}

func (c *Client) withAddrRw(addr net.Addr, fn func(*bufio.ReadWriter) error) (err error) {
	cn, err := c.getConn(addr)

	if err != nil {
		return err
	}

	defer cn.condRelease(&err)

	return fn(cn.rw)
}

func (c *Client) withKeyRw(key string, fn func(*bufio.ReadWriter) error) error {
	return c.withKeyAddr(key, func(addr net.Addr) error {
		return c.withAddrRw(addr, fn)
	})
}

func (c *Client) getFromAddr(addr net.Addr, keys []string, cb func(*Item)) error {

	return c.withAddrRw(addr, func(rw *bufio.ReadWriter) error {

		if _, err := fmt.Fprintf(rw, "gets %s\r\n", strings.Join(keys, " ")); err != nil {
			return err
		}

		if err := rw.Flush(); err != nil {
			return err
		}

		if err := parseGetResponse(rw.Reader, cb); err != nil {
			return err
		}

		return nil
	})
}

func (c *Client) flushAllFromAddr(addr net.Addr) error {
	return c.withAddrRw(addr, func(rw *bufio.ReadWriter) error {

		if _, err := fmt.Fprintf(rw, "flush_all\r\n"); err != nil {
			return err
		}

		if err := rw.Flush(); err != nil {
			return err
		}

		line, err := rw.ReadSlice('\n')
		if err != nil {
			return err
		}

		switch {
		case bytes.Equal(line, resultOk):
			break
		default:
			return fmt.Errorf("memcache: unexpected response line from flush_all: %q", string(line))
		}

		return nil
	})
}

func (c *Client) ping(addr net.Addr) error {

	return c.withAddrRw(addr, func(rw *bufio.ReadWriter) error {
		if _, err := fmt.Fprintf(rw, "version\r\n"); err != nil {
			return err
		}

		if err := rw.Flush(); err != nil {
			return err
		}

		line, err := rw.ReadSlice('\n')
		if err != nil {
			return err
		}

		switch {
		case bytes.HasPrefix(line, versionPrefix):
			break
		default:
			return fmt.Errorf("memcache: unexpected response line from ping: %q", string(line))
		}

		return nil
	})
}

func (c *Client) touchFromAddr(addr net.Addr, keys []string, expiration int32) error {
	return c.withAddrRw(addr, func(rw *bufio.ReadWriter) error {
		for _, key := range keys {
			if _, err := fmt.Fprintf(rw, "touch %s %d\r\n", key, expiration); err != nil {
				return err
			}

			if err := rw.Flush(); err != nil {
				return err
			}

			line, err := rw.ReadSlice('\n')
			if err != nil {
				return err
			}

			switch {
			case bytes.Equal(line, resultTouched):
				return nil
			case bytes.Equal(line, resultNotFound):
				return ErrCacheMiss
			default:
				return fmt.Errorf("memcache: unexpected response line from touch: %q", string(line))
			}
		}
		return nil
	})
}

func (c *Client) GetMulti(keys []string) (map[string]*Item, error) {

	var lk sync.Mutex

	m := make(map[string]*Item)

	addItemToMap := func(it *Item) {
		lk.Lock()
		defer lk.Unlock()
		m[it.Key] = it
	}

	keyMap := make(map[net.Addr][]string)

	for _, key := range keys {
		if !legalKey(key) {
			return nil, ErrMalformedKey
		}
		addr, err := c.selector.PickServer(key)
		if err != nil {
			return nil, err
		}
		keyMap[addr] = append(keyMap[addr], key)
	}

	ch := make(chan error, buffered)

	for addr, keys := range keyMap {
		go func(addr net.Addr, keys []string) {
			ch <- c.getFromAddr(addr, keys, addItemToMap)
		}(addr, keys)
	}

	var err error

	for range keyMap {
		if ge := <-ch; ge != nil {
			err = ge
		}
	}

	return m, err
}

func parseGetResponse(r *bufio.Reader, cb func(*Item)) error {
	for {
		line, err := r.ReadSlice('\n')

		if err != nil {
			return err
		}

		if bytes.Equal(line, resultEnd) {
			return nil
		}

		it := new(Item)
		size, err := scanGetResponseLine(line, it)
		if err != nil {
			return err
		}

		it.Value = make([]byte, size+2)

		_, err = io.ReadFull(r, it.Value)
		if err != nil {
			it.Value = nil
			return err
		}

		if !bytes.HasSuffix(it.Value, crlf) {
			it.Value = nil
			return fmt.Errorf("memcache: corrupt get result read")
		}

		it.Value = it.Value[:size]

		cb(it)
	}
}

func scanGetResponseLine(line []byte, it *Item) (size int, err error) {

	pattern := "VALUE %s %d %d %d\r\n"

	dest := []interface{}{&it.Key, &it.Flags, &size, &it.CasID}

	if bytes.Count(line, space) == 3 {
		pattern = "VALUE %s %d %d\r\n"
		dest = dest[:3]
	}

	n, err := fmt.Sscanf(string(line), pattern, dest...)
	if err != nil || n != len(dest) {
		return -1, fmt.Errorf("memcache: unexpected line in get response: %q", line)
	}

	return size, nil
}

func (c *Client) Set(item *Item) error {
	return c.onItem(item, (*Client).set)
}

func (c *Client) set(rw *bufio.ReadWriter, item *Item) error {
	return c.populateOne(rw, "set", item)
}

func (c *Client) Add(item *Item) error {
	return c.onItem(item, (*Client).add)
}

func (c *Client) add(rw *bufio.ReadWriter, item *Item) error {
	return c.populateOne(rw, "add", item)
}

func (c *Client) Replace(item *Item) error {
	return c.onItem(item, (*Client).replace)
}

func (c *Client) replace(rw *bufio.ReadWriter, item *Item) error {
	return c.populateOne(rw, "replace", item)
}

func (c *Client) Append(item *Item) error {
	return c.onItem(item, (*Client).append)
}

func (c *Client) append(rw *bufio.ReadWriter, item *Item) error {
	return c.populateOne(rw, "append", item)
}

func (c *Client) Prepend(item *Item) error {
	return c.onItem(item, (*Client).prepend)
}

func (c *Client) prepend(rw *bufio.ReadWriter, item *Item) error {
	return c.populateOne(rw, "prepend", item)
}

func (c *Client) CompareAndSwap(item *Item) error {
	return c.onItem(item, (*Client).cas)
}

func (c *Client) cas(rw *bufio.ReadWriter, item *Item) error {
	return c.populateOne(rw, "cas", item)
}

func (c *Client) populateOne(rw *bufio.ReadWriter, verb string, item *Item) error {

	if !legalKey(item.Key) {
		return ErrMalformedKey
	}

	var err error

	if verb == "cas" {
		_, err = fmt.Fprintf(rw, "%s %s %d %d %d %d\r\n",
			verb, item.Key, item.Flags, item.Expiration, len(item.Value), item.CasID)
	} else {
		_, err = fmt.Fprintf(rw, "%s %s %d %d %d\r\n",
			verb, item.Key, item.Flags, item.Expiration, len(item.Value))
	}

	if err != nil {
		return err
	}

	if _, err = rw.Write(item.Value); err != nil {
		return err
	}

	if _, err := rw.Write(crlf); err != nil {
		return err
	}

	if err := rw.Flush(); err != nil {
		return err
	}

	line, err := rw.ReadSlice('\n')
	if err != nil {
		return err
	}

	switch {
	case bytes.Equal(line, resultStored):
		return nil
	case bytes.Equal(line, resultNotStored):
		return ErrNotStored
	case bytes.Equal(line, resultExists):
		return ErrCASConflict
	case bytes.Equal(line, resultNotFound):
		return ErrCacheMiss
	}

	return fmt.Errorf("memcache: unexpected response line from %q: %q", verb, string(line))
}

func writeReadLine(rw *bufio.ReadWriter, format string, args ...interface{}) ([]byte, error) {

	_, err := fmt.Fprintf(rw, format, args...)

	if err != nil {
		return nil, err
	}

	if err := rw.Flush(); err != nil {
		return nil, err
	}

	return rw.ReadSlice('\n')
}

func writeExpectf(rw *bufio.ReadWriter, expect []byte, format string, args ...interface{}) error {

	line, err := writeReadLine(rw, format, args...)

	if err != nil {
		return err
	}

	switch {
	case bytes.Equal(line, resultOK):
		return nil
	case bytes.Equal(line, expect):
		return nil
	case bytes.Equal(line, resultNotStored):
		return ErrNotStored
	case bytes.Equal(line, resultExists):
		return ErrCASConflict
	case bytes.Equal(line, resultNotFound):
		return ErrCacheMiss
	}

	return fmt.Errorf("memcache: unexpected response line: %q", string(line))
}

func (c *Client) Delete(key string) error {
	return c.withKeyRw(key, func(rw *bufio.ReadWriter) error {
		return writeExpectf(rw, resultDeleted, "delete %s\r\n", key)
	})
}

func (c *Client) DeleteAll() error {
	return c.withKeyRw("", func(rw *bufio.ReadWriter) error {
		return writeExpectf(rw, resultDeleted, "flush_all\r\n")
	})
}

func (c *Client) Ping() error {
	return c.selector.Each(c.ping)
}

func (c *Client) Increment(key string, delta uint64) (newValue uint64, err error) {
	return c.incrDecr("incr", key, delta)
}

func (c *Client) Decrement(key string, delta uint64) (newValue uint64, err error) {
	return c.incrDecr("decr", key, delta)
}

func (c *Client) incrDecr(verb, key string, delta uint64) (uint64, error) {
	var val uint64
	err := c.withKeyRw(key, func(rw *bufio.ReadWriter) error {

		line, err := writeReadLine(rw, "%s %s %d\r\n", verb, key, delta)

		if err != nil {
			return err
		}

		switch {
		case bytes.Equal(line, resultNotFound):
			return ErrCacheMiss
		case bytes.HasPrefix(line, resultClientErrorPrefix):
			errMsg := line[len(resultClientErrorPrefix) : len(line)-2]
			return errors.New("memcache: client error: " + string(errMsg))
		}

		val, err = strconv.ParseUint(string(line[:len(line)-2]), 10, 64)
		if err != nil {
			return err
		}

		return nil
	})

	return val, err
}

func (c *Client) Close() error {

	c.mutx.Lock()
	defer c.mutx.Unlock()

	var ret error

	for _, conns := range c.freeconn {
		for _, c := range conns {
			if err := c.nc.Close(); err != nil && ret == nil {
				ret = err
			}
		}
	}

	c.freeconn = nil

	return ret
}
