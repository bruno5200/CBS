package client

import (
	"crypto/tls"
	"net/http"
)

type client struct{ client http.Client }

func NewClient() *client {
	return &client{client: http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}}
}
