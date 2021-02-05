package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tpmanc/go-projects/db"
	"github.com/tpmanc/go-projects/handlers"
	"github.com/tpmanc/go-projects/middlewares"
	"log"
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

	r.HandleFunc("/projects", handlers.ProjectsHandler).Methods("GET")
	r.HandleFunc("/project/{id:[0-9]+}", handlers.ProjectHandler).Methods("GET")
	r.HandleFunc("/project/save", handlers.ProjectSaveHandler).Methods("POST")
	r.HandleFunc("/project/delete", handlers.ProjectDeleteHandler).Methods("DELETE")

	log.Println("Start server")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}