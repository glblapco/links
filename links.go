package main

import (
	"log"
	"net/http"
	"proj/links/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
