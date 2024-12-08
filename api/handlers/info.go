package handlers

import (
	"api/database"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    db, err := database.Connect()

    if err != nil {
        json.NewEncoder(w).Encode("Error connecting to database")
        return 
    }

    url, err := database.GetUrl(db, vars["url"])

    if err != nil {
        json.NewEncoder(w).Encode("Error getting url")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(url)
}
