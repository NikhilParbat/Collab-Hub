package main

import (
	"log"

	"github.com/NikhilParbat/Collab-Hub/db"
	"github.com/NikhilParbat/Collab-Hub/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/projects", handlers.CreateProject)
	r.POST("/projects/:projectId/join/:userId", handlers.JoinProject)

	log.Fatal(r.Run(":8080"))
}
