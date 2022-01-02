package handlers

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"proj/links/models"
	"proj/links/repository"
)

func list(rw http.ResponseWriter, req *http.Request) {
	links := repository.GetLinks()
	json, _ := json.Marshal(links)

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(json)

	log.Println("Response: ", 200)
}

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
