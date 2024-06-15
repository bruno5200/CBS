package env

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var MuDB = sync.Mutex{}

// environmentDB contains the environment variables for the postgres databases
type environmentDB struct {
	dbHost   string
	dbHostOK bool
	dbPort   string
	dbPortOK bool
	dbUser   string
	dbUserOK bool
	dbPass   string
	dbPassOK bool
	dbName   string
	dbNameOK bool
	dbSSL    string
	dbSSLOK  bool
}

// InitDB initializes the environment variables for the postgres databases
//
// if the variables are not set, the application will panic
func (e *environment) initDB(wg *sync.WaitGroup) {
	MuDB.Lock()
	e.dbHost, e.dbHostOK = os.LookupEnv("DB_HOST")
	e.dbPort, e.dbPortOK = os.LookupEnv("DB_PORT")
	e.dbUser, e.dbUserOK = os.LookupEnv("DB_USER")
	e.dbPass, e.dbPassOK = os.LookupEnv("DB_PASSWORD")
	e.dbName, e.dbNameOK = os.LookupEnv("DB_NAME")
	e.dbSSL, e.dbSSLOK = os.LookupEnv("DB_SSL")
	MuDB.Unlock()
	wg.Done()
}

// `DB_HOST`: Host de la base de datos.
//
// if Host is not set, the application will panic
func (e *environment) GetDBHost() (val string) {
	if e.dbHostOK {
		val = e.dbHost
	} else {
		log.Panic("PostgreSQL: DB_HOST not set")
	}
	return
}

// `DB_PORT`: Puerto de la base de datos.
//
// if Port is not set, the application will panic
func (e *environment) GetDBPort() (val int) {
	if e.dbPortOK {
		value, err := strconv.Atoi(e.dbPort)

		if err != nil {
			log.Printf("Env: Invalid BD_PORT %s, err: %s", e.dbPort, err.Error())
			return
		}
		val = value
	} else {
		log.Panic("PostgreSQL: DB_PORT not set")
	}
	return
}

// `DB_USER`: Usuario de la base de datos.
//
// if User is not set, the application will panic
func (e *environment) GetDBUser() (val string) {
	if e.dbUserOK {
		val = e.dbUser
	} else {
		log.Panic("PostgreSQL: DB_USER not set")
	}
	return
}

// `DB_PASSWORD`: Contraseña de la base de datos.
//
// if Password is not set, the application will panic
func (e *environment) GetDBPass() (val string) {
	if e.dbPassOK {
		val = e.dbPass
	} else {
		log.Panic("PostgreSQL: DB_PASS not set")
	}
	return
}

// `DB_NAME`: Nombre de la base de datos.
//
// if Name is not set, the application will panic
func (e *environment) GetDBName() (val string) {
	if e.dbNameOK {
		val = e.dbName
	} else {
		log.Panic("PostgreSQL: DB_NAME not set")
	}
	return
}

// `DB_SSL`: SSL de la base de datos.
//
// Por defecto es `disable`.
func (e *environment) GetDBSSL() (val string) {
	if e.dbSSLOK {
		val = e.dbSSL
	} else {
		val = "disable"
	}
	return
}
