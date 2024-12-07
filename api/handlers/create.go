package handlers

import (
	"encoding/json"
	"net/http"

    "api/database"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    short := r.FormValue("short")   
    long := r.FormValue("long")

    if short == "" || long == "" {
        json.NewEncoder(w).Encode("Short and long URL required")
        return
    }

    db, err := database.Connect()

    if err != nil {
        json.NewEncoder(w).Encode("Error connecting to database")
        return
    }

    err = database.CreateUrl(db, short, long)

    if err != nil {
        json.NewEncoder(w).Encode("Error creating URL")
        return
    }

    json.NewEncoder(w).Encode("URL created successfully")
}
