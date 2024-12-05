package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})

	http.ListenAndServe(":3000", nil)
}
