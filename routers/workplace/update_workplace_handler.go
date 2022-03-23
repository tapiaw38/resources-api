package workplace

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/resources-api/models"

	workplace "github.com/tapiaw38/resources-api/database/workplace"
)

func UpdateWorkplaceHandler(w http.ResponseWriter, r *http.Request) {

	var wp models.Workplace

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&wp)

	// Name and code are required
	if wp.Name == "" || wp.Code == "" {
		http.Error(w, "An error occurred, name, code are required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a workplace "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	workplace, err := workplace.UpdateWorkplace(ctx, id, wp)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a workplace in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workplace)
}
