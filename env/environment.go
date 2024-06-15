package env

import "sync"

var (
	e    *environment = &environment{}
	once sync.Once
	wg   sync.WaitGroup
)

// Centralizer for environment variables
type environment struct {
	environmentApp
	environmentDB
	environmentCache
	environmentBlob
}

// Init initializes the environment variables
func Init() {
	once.Do(func() {
		e.init()
	})
}

func (e *environment) init() {
	wg.Add(4)
	go e.initApp(&wg)
	go e.initDB(&wg)
	go e.initCache(&wg)
	go e.initBlob(&wg)
	wg.Wait()
}

func Env() *environment {
	return e
}
