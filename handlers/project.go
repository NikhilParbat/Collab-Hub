package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	firebaseapp "github.com/NikhilParbat/Collab-Hub/firebase"
	"github.com/NikhilParbat/Collab-Hub/middleware"
	"github.com/NikhilParbat/Collab-Hub/models"
	"github.com/gorilla/mux"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(middleware.UserIDKey).(string)

	var project models.Project
	json.NewDecoder(r.Body).Decode(&project)

	project.OwnerID = uid
	project.Status = "open"
	project.CreatedAt = time.Now()

	ref, _, err := firebaseapp.Firestore.
		Collection("projects").
		Add(context.Background(), project)

	if err != nil {
		http.Error(w, "Failed to create project", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"projectId": ref.ID,
	})
}

func JoinProject(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(middleware.UserIDKey).(string)
	projectId := mux.Vars(r)["id"]

	_, err := firebaseapp.Firestore.
		Collection("projects").
		Doc(projectId).
		Collection("requests").
		Doc(uid).
		Set(context.Background(), map[string]interface{}{
			"requestedAt": time.Now(),
		})

	if err != nil {
		http.Error(w, "Failed to request join", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "request sent"})
}
