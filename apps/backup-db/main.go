package main

import (
	"github.com/gorilla/mux"
	"github.com/tpmanc/backup-db/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/run", handlers.RunHandler).Methods("POST")

	log.Println("Start server")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
