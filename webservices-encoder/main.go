package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0] 
		}

		m := map[string]string{"name" : name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})

	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal(err)
	}
} 