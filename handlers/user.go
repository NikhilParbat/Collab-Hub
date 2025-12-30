package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	firebaseapp "github.com/NikhilParbat/Collab-Hub/firebase"
	"github.com/NikhilParbat/Collab-Hub/middleware"
	"github.com/NikhilParbat/Collab-Hub/models"
)

func InitUser(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(middleware.UserIDKey).(string)

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.CreatedAt = time.Now()

	_, err := firebaseapp.Firestore.Collection("users").Doc(uid).Set(
		context.Background(), user,
	)

	if err != nil {
		http.Error(w, "Failed to create user", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "user created"})
}
