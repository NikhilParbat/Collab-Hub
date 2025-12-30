package models

import "time"

type Project struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	OwnerID        string    `json:"ownerId"`
	SkillsRequired []string  `json:"skillsRequired"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
}
