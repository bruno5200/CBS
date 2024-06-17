package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
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
