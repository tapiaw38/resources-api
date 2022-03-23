package workplace

import (
	"encoding/json"
	"net/http"

	workplace "github.com/tapiaw38/resources-api/database/workplace"
	"github.com/tapiaw38/resources-api/models"
)

// CreateWorkplaceHandler is a handler for creating a new workplace
func CreateWorkplaceHandler(w http.ResponseWriter, r *http.Request) {

	var wp models.Workplace

	err := json.NewDecoder(r.Body).Decode(&wp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a workplace "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	workplace, err := workplace.CreateWorkplace(ctx, &wp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a workplace in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workplace)

}
