package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	user "github.com/tapiaw38/resources-api/database/user"
)

// GetUsersHandler handles the request to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	users, err := user.GetUsers(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get users "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

// GetUserByIdHandler handles the request to get a user by id
func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	user, err := user.GetUserById(ctx, id)

	if err != nil {
		http.Error(w, "No record found with that id "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

// GetUserByUsernameHandler handles the request to get a user by username
func GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	username := r.URL.Query().Get("username")

	user, err := user.GetUserByUsername(ctx, username)

	if err != nil {
		http.Error(w, "No record found with that username "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
