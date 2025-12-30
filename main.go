package main

import (
	"log"
	"net/http"

	firebaseapp "github.com/NikhilParbat/Collab-Hub/firebase"
	"github.com/NikhilParbat/Collab-Hub/handlers"
	"github.com/NikhilParbat/Collab-Hub/middleware"
	"github.com/gorilla/mux"
	// "github.com/labstack/echo/v4"
)

func main() {
	firebaseapp.InitFirebase()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Collab-Hub!"))
	}).Methods("GET")
	api.HandleFunc("/users/init", handlers.InitUser).Methods("POST")
	api.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	api.HandleFunc("/projects/{id}/join", handlers.JoinProject).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
