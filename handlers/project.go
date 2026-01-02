package handlers

import (
	"strconv"

	"github.com/NikhilParbat/Collab-Hub/db"

	"github.com/gin-gonic/gin"
)

// Create project
func CreateProject(c *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		OwnerID     int    `json:"ownerId"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	var projectID int
	err := db.DB.QueryRow(`
		INSERT INTO projects (title, description, owner_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`, body.Title, body.Description, body.OwnerID).Scan(&projectID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"projectId": projectID})
}

// Join project
func JoinProject(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("projectId"))
	userID, _ := strconv.Atoi(c.Param("userId"))

	_, err := db.DB.Exec(`
		INSERT INTO project_members (project_id, user_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, projectID, userID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "joined project"})
}
