package main

import (
	"log"
	"os"

	db "github.com/NikhilParbat/Collab-Hub/db"
	handlers "github.com/NikhilParbat/Collab-Hub/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Initializing DB...")
	db.InitDB()

	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/projects", handlers.CreateProject)
	r.POST("/projects/:projectId/join/:userId", handlers.JoinProject)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on port", port)
	log.Fatal(r.Run(":" + port))
}
