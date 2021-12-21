package utils

import (
	"log"
	"net/http"
)

type Playlist struct {
	Videos []string
	Name   string
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleResponse(res *http.Response) {
	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Fatal("No se ha podido conectar con la API por la siguiente razon: ")
	}
}
