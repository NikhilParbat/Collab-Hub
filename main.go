package main

import (
	"log"
	"os"

	db "github.com/NikhilParbat/Collab-Hub/db/sqlc"
	"github.com/NikhilParbat/Collab-Hub/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	log.Println("Initializing DB...")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}
	conn := db.InitDB(dsn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := server.NewServer(conn)
	log.Printf("Starting server on port %s...", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
