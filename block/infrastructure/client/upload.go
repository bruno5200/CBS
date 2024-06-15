package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (c *client) UploadToBlob(content, url, key, name string) (err error) {

	decodedData, err := base64.StdEncoding.DecodeString(content)

	if err != nil {
		log.Println("Error al decodificar Base64:", err)
		return
	}

	var requestBody bytes.Buffer

	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile(key, name)

	if err != nil {
		log.Println("Error al crear el formulario:", err)
		return
	}

	part.Write(decodedData)

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, url, &requestBody)

	if err != nil {
		log.Println("Error al crear la solicitud HTTP:", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(req)

	if err != nil {
		log.Println("Error al enviar la solicitud HTTP:", err)
		return
	}

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		log.Printf("error uploading document to blob storage, status: %s", resp.Status)
		err = fmt.Errorf("error uploading document to blob storage, status: %s", resp.Status)
		return
	}

	return
}

func (c *client) DownloadFromBlob(url string) (content string, err error) {

	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error al descargar el archivo:", err)
		return
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	
	buf.ReadFrom(resp.Body)

	content = buf.String()

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
