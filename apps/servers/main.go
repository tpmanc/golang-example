package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tpmanc/servers/db"
	"github.com/tpmanc/servers/handlers"
	"github.com/tpmanc/servers/middlewares"
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

	r.HandleFunc("/servers", handlers.ServersHandler).Methods("GET")
	r.HandleFunc("/server/{id:[0-9]+}", handlers.ServerHandler).Methods("GET")
	r.HandleFunc("/server/save", handlers.ServerSaveHandler).Methods("POST")
	r.HandleFunc("/server/delete", handlers.ServerDeleteHandler).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
