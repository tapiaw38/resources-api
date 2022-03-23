package workplace

import (
	"encoding/json"
	"net/http"

	workplace "github.com/tapiaw38/resources-api/database/workplace"
)

// GetWorkplacesHandler is a handler for getting a workplace
func GetWorkplacesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	workplaces, err := workplace.GetWorkplaces(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get workplaces "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workplaces)

}
