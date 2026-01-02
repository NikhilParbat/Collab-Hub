package handlers

import (
	"github.com/NikhilParbat/Collab-Hub/db"

	"github.com/gin-gonic/gin"
)

// Create user
func CreateUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	var id int
	err := db.DB.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id
	`, body.Name, body.Email).Scan(&id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"userId": id})
}

// Delete user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "user deleted"})
}
