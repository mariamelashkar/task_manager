package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"task/internal/redis"
	"task/models"
)
var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, valid, err := CheckUserCredentials(loginRequest.Username, loginRequest.Password)
	if err != nil || !valid {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	userRole, err := GetUserRole(loginRequest.Username)
	if err != nil {
		http.Error(w, "Could not retrieve user role", http.StatusInternalServerError)
		return
	}

	tokenString, err := redis.GenerateToken(strconv.Itoa(userID), userRole)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	err = redis.StoreToken(tokenString)
	if err != nil {
		http.Error(w, "Could not store token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}