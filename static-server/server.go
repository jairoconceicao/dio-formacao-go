package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting on :8080")

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
