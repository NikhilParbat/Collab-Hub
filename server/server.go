package server

import (
	"log"

	db "github.com/NikhilParbat/Collab-Hub/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router *gin.Engine
	Store  db.Store
}

func NewServer(conn *pgxpool.Pool) *Server {
	// Initialize DB pool (pgxpool)

	if conn == nil {
		log.Fatal("db pool not initialized")
	}

	// Create store using pgx pool
	store := db.NewStore(conn)

	s := &Server{
		router: gin.Default(),
		Store:  store,
	}

	s.setupRouter()

	return s
}

func (s *Server) setupRouter() {

	r := gin.Default()

	r.POST("/users", s.CreateUser)

	s.router = r
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
