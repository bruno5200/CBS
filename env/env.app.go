package env

import (
	"os"
	"sync"
)

var (
	MuApp = sync.Mutex{}
)

// environmentApp contains the environment variables for the application
type environmentApp struct {
	port     string
	portOK   bool
	url      string
	urlOK    bool
	secrue   string
	secureOK bool
}

// InitApp initializes the environment variables for the application
//
// if the variables are not set default values will be used
func (e *environment) initApp(wg *sync.WaitGroup) {
	MuApp.Lock()
	e.port, e.portOK = os.LookupEnv("PORT")
	e.url, e.urlOK = os.LookupEnv("URL")
	e.secrue, e.secureOK = os.LookupEnv("SECURE")
	MuApp.Unlock()
	wg.Done()
}

// `PORT`: Puerto en el que se ejecutará el servidor.
//
// Por defecto es `8080`.
func (e *environment) GetPort() (val string) {
	if !e.portOK {
		val = "8080"
	} else {
		val = e.port
	}
	return
}

// `URL`: Dirección url con la que se envia los correos.
//
// Por defecto es `https://dev-cbs-2d7ukhlbva-uc.a.run.app`.
func (e *environment) GetUrl() (val string) {
	if !e.urlOK {
		val = "https://dev-cbs-2d7ukhlbva-uc.a.run.app"
	} else {
		val = e.url
	}
	return
}

// `SECURE`: habilita la encriptación de los datos mediante TLS.
//
// Por defecto es `false`.
func (e *environment) GetSecure() (val bool) {
	if e.secureOK {
		val = e.secrue == "true"
	}
	return
}
