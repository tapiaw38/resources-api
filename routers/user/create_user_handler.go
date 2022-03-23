package user

import (
	"encoding/json"
	"net/http"

	user "github.com/tapiaw38/resources-api/database/user"
	"github.com/tapiaw38/resources-api/models"
)

// CreateUserHandler handles the request to create a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := user.CreateUser(ctx, &u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user in database "+err.Error(), 400)
		return
	}

	u.PasswordHash = ""

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
