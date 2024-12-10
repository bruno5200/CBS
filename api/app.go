package main

import (
	"log"

	"github.com/bruno5200/CSM/memcache"
)

func main() {
	mc := memcache.New( /*"127.0.0.1:11212", "127.0.0.1:11213",  "34.176.89.192:11211",*/ "34.176.110.27:11211")

	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my first value in memecached")})

	it, err := mc.Get("foo")
	if err != nil {
		println(err.Error())
	}

	// DEV
	// if err := mc.Delete("8089e24d-0c07-4dbe-ae50-459c35891f35"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("4e2166b9-7ed0-4d65-ad20-da4fa8ef64fb"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("5286d088-8d63-4ae5-85db-a55d86b10856"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("7399a531-0463-409b-9480-8839b4de42b4"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("941dd882-9622-4c64-9865-c40374ed2ed7"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("64ea7bc7-30ba-41e5-b327-bcae7128312b"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("40794b4f-a873-4703-a5cd-93771a9370d5"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("42fa80ec-8d2e-4178-aecf-ecbc3d4cb29f"); err != nil {log.Printf("MEM %s", err)}
	// if err := mc.Delete("771fe73b-d8e9-4a3f-a837-1bd6bef3b688"); err != nil {log.Printf("MEM %s", err)}

	// CERT
	if err := mc.Delete("8089e24d-0c07-4dbe-ae50-459c35891f35"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("4e2166b9-7ed0-4d65-ad20-da4fa8ef64fb"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("5286d088-8d63-4ae5-85db-a55d86b10856"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("7399a531-0463-409b-9480-8839b4de42b4"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("941dd882-9622-4c64-9865-c40374ed2ed7"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("6542b4e0-2c13-4431-bd1d-8a95c5471917"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("809747db-dce0-446c-9263-617cabcfa204"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("5996e3e4-8044-4e86-bcde-10e30ac005d7"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("75e7680b-910f-45f7-9a38-637b49a4d492"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("a0b4cbdd-fe9f-4c69-90a0-9739a251fc0f"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("42fa80ec-8d2e-4178-aecf-ecbc3d4cb29f"); err != nil {log.Printf("MEM %s", err)}
	if err := mc.Delete("5394f9c0-8160-4f93-9b64-86c8b4aebc9d"); err != nil {log.Printf("MEM %s", err)}

	println(string(it.Value))

	mc.Close()
}
