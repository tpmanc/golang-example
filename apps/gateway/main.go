package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tpmanc/gateway/db"
	"github.com/tpmanc/gateway/handlers"
	"github.com/tpmanc/gateway/middleware"
	"net/http"
)

func main() {
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		errors.New("cant find .env file")
	}

	db.Init(env)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	// router.HandleFunc("/logout", sugnupHandler).Methods("POST")

	//router.Handle("/projects", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProjectsHandler))).Methods("GET")
	//router.Handle("/project/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProjectHandler))).Methods("GET")
	//router.Handle("/project/save", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProjectSaveHandler))).Methods("POST")

	router.Use(middleware.LogMiddleware)
	
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8000", nil)
}