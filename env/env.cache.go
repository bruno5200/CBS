package env

import (
	"os"
	"strings"
	"sync"
)

var MuCache = &sync.Mutex{}

// environmentCache contains the environment variables for memcached connection
type environmentCache struct {
	cacheURLs   string
	cacheURLsOK bool
}

// InitCache initializes the environment variables for memcached connection
//
// if the variables are not set default values will be used
func (e *environment) initCache(wg *sync.WaitGroup) {
	MuCache.Lock()
	e.cacheURLs, e.cacheURLsOK = os.LookupEnv("CACHE_URLS")
	MuCache.Unlock()
	wg.Done()
}

// `CACHE_URLS`: URL de conexión a memcached.
//
// Por defecto es `[localhost:11211]`.
func (e *environment) GetCacheURLs() (val []string) {
	if !e.cacheURLsOK {
		val = []string{"localhost:11211"}
	} else if strings.Contains(e.cacheURLs, ",") {
		val = strings.Split(e.cacheURLs, ",")
	} else {
		val = []string{e.cacheURLs}
	}
	return
}
