package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.RedirectHandler)

    urlRouter := router.PathPrefix("/api/url").Subrouter()

    urlRouter.HandleFunc("/{id}", handlers.InfoHandler).Methods(http.MethodGet)
    urlRouter.HandleFunc("/", handlers.ListHandler).Methods(http.MethodGet)
    urlRouter.HandleFunc("/", handlers.CreateHandler).Methods(http.MethodPost)
	urlRouter.HandleFunc("/", handlers.UpdateHandler).Methods(http.MethodPatch)
    urlRouter.HandleFunc("/", handlers.DeleteHandler).Methods(http.MethodDelete) 

    fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
