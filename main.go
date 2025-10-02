package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to Auth"))
		if err != nil {
			log.Printf("Error writing response: %s\n", err.Error())
		}
	})

	port := "8080"
	log.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
