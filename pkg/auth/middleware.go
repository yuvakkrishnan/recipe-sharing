package auth

import (
	"context"
	"net/http"

	"github.com/yuvakkrishnan/user-service/pkg/response"
)

type contextKey string

const userKey contextKey = "userID"

// Middleware validates JWT token and extracts user ID
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			response.Error(w, http.StatusUnauthorized, "Authorization token missing")
			return
		}

		userID, err := ValidateToken(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Store userID in context for use in handlers
		ctx := context.WithValue(r.Context(), userKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserIDFromContext extracts the user ID from the request context
func UserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(userKey).(int64)
	return userID, ok
}
