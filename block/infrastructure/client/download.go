package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (c *client) DownloadFromBlob(url string) (content string, err error) {

	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error al descargar el archivo:", err)
		return
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)

	buf.ReadFrom(resp.Body)

	content = base64.StdEncoding.EncodeToString(buf.Bytes())

	return
}

func (c *client) DownloadFromBlobStorage(url, path string) (err error) {

	res, err := c.client.Get(url)

	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		log.Printf("error downloading file from blob storage, status: %s", res.Status)
		err = fmt.Errorf("error downloading file from blob storage, status: %s", res.Status)
		return
	}

	file, err := os.Create(path)

	if err != nil {
		log.Printf("Error creating file: %s", err)
		return
	}

	defer file.Close()

	if _, err = io.Copy(file, res.Body); err != nil {
		log.Printf("Error copying file: %s", err)
		return
	}

	return
}
