package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/resources-api/models/workplace"
)

type WorkplaceRouter struct {
	Storage workplace.Storage
}

// CreateWorkplaceHandler is a handler for creating a new workplace
func (wps *WorkplaceRouter) CreateWorkplaceHandler(w http.ResponseWriter, r *http.Request) {

	var wp workplace.Workplace

	err := json.NewDecoder(r.Body).Decode(&wp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a workplace "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	workplace, err := wps.Storage.CreateWorkplace(ctx, &wp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter a workplace in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workplace)

}

// GetWorkplacesHandler is a handler for getting a workplace
func (wps *WorkplaceRouter) GetWorkplacesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	workplaces, err := wps.Storage.GetWorkplaces(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get workplaces "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workplaces)

}

// UpdateWorkplaceHandler handles the request to update a workplace
func (wps *WorkplaceRouter) UpdateWorkplaceHandler(w http.ResponseWriter, r *http.Request) {

	var wp workplace.Workplace

	id := mux.Vars(r)["id"]

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
	workplace, err := wps.Storage.UpdateWorkplace(ctx, id, wp)

	if err != nil {
		http.Error(w, "An error occurred when trying to update a workplace in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workplace)
}

func (wps *WorkplaceRouter) DeleteWorkplaceHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err := wps.Storage.DeleteWorkplace(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred while trying to delete a record from the database"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
