package middleware

import (
	"context"
	"net/http"
	"strings"

	firebaseapp "github.com/NikhilParbat/Collab-Hub/firebase"
)

type contextKey string

const UserIDKey contextKey = "userId"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Replace(header, "Bearer ", "", 1)
		token, err := firebaseapp.Auth.VerifyIDToken(context.Background(), tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
