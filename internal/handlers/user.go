package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuvakkrishnan/user-service/internal/models"
	"github.com/yuvakkrishnan/user-service/pkg/auth"
	"github.com/yuvakkrishnan/user-service/pkg/response"
)

type UserService interface {
	Register(user *models.User) error
	Login(username, password string) (string, error)
	GetProfile(userID int64) (*models.User, error)
	ForgotPassword(email string) error
	ResetPassword(token, newPassword string) error
}

func Register(svc UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Entering Register")

		var user models.User
		// Parse and decode the request body into the User struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Validate the user input
		if user.Username == "" || user.Password == "" || user.Email == "" {
			response.Error(w, http.StatusBadRequest, "Missing required fields")
			return
		}

		// Call the user service to register the user
		err = svc.Register(&user)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to register user")
			return
		}

		// Return success response
		response.JSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
	}
}
func Login(svc UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Entering Login")
		// Parse and validate input
		// Call user service to authenticate user
		// Return JWT token if successful
	}
}

func GetProfile(svc UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Entering GetProfile")

		userID, ok := auth.UserIDFromContext(r.Context())
		if !ok {
			response.Error(w, http.StatusUnauthorized, "User ID not found in context")
			return
		}

		user, err := svc.GetProfile(userID)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Error fetching profile")
			return
		}

		response.JSON(w, http.StatusOK, user)
	}
}

func ForgotPassword(svc UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle forgot password logic
		log.Printf("Entering ForgotPassword")

	}
}

func ResetPassword(svc UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle reset password logic
		log.Printf("Entering ResetPassword")

	}
}

func NewRouter(svc UserService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", Register(svc)).Methods("POST")
	r.HandleFunc("/login", Login(svc)).Methods("POST")
	r.Handle("/profile", auth.Middleware(GetProfile(svc))).Methods("GET") // Changed to r.Handle
	r.HandleFunc("/forgot-password", ForgotPassword(svc)).Methods("POST")
	r.HandleFunc("/reset-password", ResetPassword(svc)).Methods("POST")
	return r
}
