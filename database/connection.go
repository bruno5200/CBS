package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/bruno5200/CSM/env"
	_"github.com/lib/pq"
)

var (
	// db is a pointer to gorm.DB
	db           *sql.DB
	once         sync.Once
	MuPostgresDB = sync.Mutex{}

	user     string
	password string
	host     string
	name     string
	port     int
)

// Connect creates a connection to database
func NewPostgresDB() {

	once.Do(func() {

		connStr := connection()

		var err error

		if db, err = sql.Open("postgres", connStr); err != nil {
			log.Printf("Error connecting to database %s", err)
		}

		if err := db.Ping(); err != nil {
			log.Panicf("Error pinging database %s", err)
		} else {
			fmt.Println("Connected to postgres")
		}
	})
}

func connection() string {

	en := env.Env()
	env.MuDB.Lock()
	host = en.GetDBHost()
	port = en.GetDBPort()
	name = en.GetDBName()
	user = en.GetDBUser()
	password = en.GetDBPass()
	env.MuDB.Unlock()

	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, name, port)
}

func PostgresDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
