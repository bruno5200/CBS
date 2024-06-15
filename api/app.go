package main

import (
	"github.com/bruno5200/CSM/memcache"
)

func main() {
	mc := memcache.New("127.0.0.1:11212","127.0.0.1:11213")
	
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my first value in memecached")})

	it, err := mc.Get("foo")
	if err != nil {
		println(err.Error())
	}
	println(string(it.Value))

	mc.Close()

}
