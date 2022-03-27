package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tapiaw38/resources-api/models"

	"github.com/tapiaw38/resources-api/jwt"

	auth "github.com/tapiaw38/resources-api/database/auth"
)

// Login handles the request to login a user
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "The email or password are invalid "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "The user email is required", 400)
		return
	}

	foundUser, exist := auth.Login(user.Email, user.Password)

	if !exist {
		http.Error(w, "The email or password are invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(foundUser)

	if err != nil {
		http.Error(w, "An error occurred while generating the token"+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		User: models.User{
			ID:        foundUser.ID,
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			Username:  foundUser.Username,
			Email:     foundUser.Email,
			Picture:   foundUser.Picture,
			IsActive:  foundUser.IsActive,
			IsAdmin:   foundUser.IsAdmin,
		},
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	//coquie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
