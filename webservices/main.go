package main

import (
	"net/http"
	"log"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal(err)
	}
}