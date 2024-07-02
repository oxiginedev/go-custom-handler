package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("something went wrong")
	}))

	log.Println("server started and listening...")
	log.Fatal(http.ListenAndServe(":6185", mux))
}
