package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tpmanc/databases/db"
	"github.com/tpmanc/databases/handlers"
	"github.com/tpmanc/databases/middlewares"
	"net/http"
)

func main() {
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		errors.New("cant find .env file")
	}

	db.Init(env)

	r := mux.NewRouter()
	r.Use(middlewares.LogMiddleware)

	r.HandleFunc("/databases", handlers.DatabasesHandler).Methods("GET")
	r.HandleFunc("/database/{id:[0-9]+}", handlers.DatabaseHandler).Methods("GET")
	r.HandleFunc("/database/save", handlers.DatabasesSaveHandler).Methods("POST")
	r.HandleFunc("/database/delete", handlers.DatabasesDeleteHandler).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
