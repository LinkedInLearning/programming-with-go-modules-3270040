package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vladimirvivien/emojiis/search"
)

func main() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		var params search.Params
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Got request: %#v", params)

		// write result
		w.Header().Set("Content-Type", "application/json")
		result := search.ByDescription(params)
		if err := json.NewEncoder(w).Encode(&result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/search/tags", func(w http.ResponseWriter, r *http.Request) {
		var tags []string
		if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Got request: %#v", tags)

		// write result
		w.Header().Set("Content-Type", "application/json")
		result := search.ByTags(tags...)
		if err := json.NewEncoder(w).Encode(&result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})


	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
