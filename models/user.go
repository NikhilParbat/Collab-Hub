package models

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Skills    []string  `json:"skills"`
	CreatedAt time.Time `json:"createdAt"`
}
