package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
