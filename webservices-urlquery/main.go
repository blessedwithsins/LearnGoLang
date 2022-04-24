package main

import (
	"net/http"
	"log"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0] 
		}
		w.Write([]byte("Hello " + name))
	})

	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal(err)
	}
}


