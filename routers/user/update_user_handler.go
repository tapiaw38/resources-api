package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	user "github.com/tapiaw38/resources-api/database/user"
	"github.com/tapiaw38/resources-api/models"
)

// UpdateUserHandler handles the request to update a user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a user "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	user, err := user.UpdateUser(ctx, id, u)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a user in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
