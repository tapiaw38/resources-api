package user

import (
	"net/http"

	user "github.com/tapiaw38/resources-api/database/user"
)

// DeleteUserHandler handles the request to delete a user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err := user.DeleteUser(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred while trying to delete a record from the database"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
