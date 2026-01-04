package server

import (
	db "github.com/NikhilParbat/Collab-Hub/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	user, err := s.Store.CreateUser(
		c.Request.Context(),
		db.CreateUserParams{
			Name:  body.Name,
			Email: body.Email,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}
