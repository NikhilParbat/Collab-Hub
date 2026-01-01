package firebaseapp

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var (
	Firestore *firestore.Client
	Auth      *auth.Client
)

func InitFirebase() {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "collab-hub-ea4da",
	})
	if err != nil {
		log.Fatal(err)
	}

	Firestore, err = app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Auth, err = app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
