package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"api/database"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    db, err := database.Connect()

    if err != nil {
        json.NewEncoder(w).Encode("Error connecting to database")
        return
    }

    err = database.DeleteUrl(db, vars["url"])

    if err != nil {
        json.NewEncoder(w).Encode("Error deleting URL")
        return
    }


    json.NewEncoder(w).Encode("URL deleted successfully")
}
