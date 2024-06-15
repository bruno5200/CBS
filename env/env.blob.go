package env

import (
	"os"
	"sync"
)

var (
	Mublob = sync.Mutex{}
)

type environmentBlob struct {
	blobUrl   string
	blobUrlOK bool
}

// Initblob initializes the environment variables for the blob
//
// if the variables are not sert default values will be used
func (e *environment) initBlob(wg *sync.WaitGroup) {
	Mublob.Lock()
	e.blobUrl, e.blobUrlOK = os.LookupEnv("BLOB_URL")
	Mublob.Unlock()
	wg.Done()
}

// `BLOB_URL`: URL del blob.
//
// Por defecto es `https://blob.gutier.lat/` + path.
func (e *environment) GetBlobUrl(path string) (val string) {
	if e.blobUrlOK {
		val = e.blobUrl + "/" + path
	} else {
		val = "https://blob.gutier.lat/" + path
	}
	return
}
