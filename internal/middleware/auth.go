package middleware

import (
	"context"
	"net/http"

	"github.com/yuvakkrishnan/user-service/pkg/auth"
	"github.com/yuvakkrishnan/user-service/pkg/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			response.JSON(w, http.StatusUnauthorized, "Authorization token missing")
			return
		}

		userID, err := auth.ValidateToken(token)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Store userID in context for use in handlers
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
