package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api/database"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
    db, err := database.Connect()

    if err != nil {
        json.NewEncoder(w).Encode("Error connecting to database")
        return
    }

    urls, err := database.GetAllUrl(db)

    fmt.Println(urls)

    if err != nil { 
        json.NewEncoder(w).Encode("Error getting URLs")
        return
    }   

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(urls)
}

