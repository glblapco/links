package handlers

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"proj/links/models"
	"proj/links/repository"
)

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request received: ". req.Method)

	switch req.Method {
		case http.MethodGet:
			list(rw, req)
			break
		case http.MethodPost:
			add(rw, req)
			break
		default:
			rw.WriteHeader(405)
			rw.Write([]byte("Method not allowed."))
			break
	}
}
