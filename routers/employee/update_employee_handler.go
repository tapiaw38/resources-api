package employee

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/resources-api/models"

	employee "github.com/tapiaw38/resources-api/database/employee"
)

// UpdateEmployeeHandler handles the request to update an employee
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	var e models.Employee

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&e)

	if e.FirstName == "" || e.LastName == "" {
		http.Error(w, "An error occurred when trying to enter an employee "+err.Error(), 400)
		return
	}

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := employee.UpdateEmployee(ctx, id, e)

	if err != nil {
		http.Error(w, "An error occurred when trying to update an employee in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}
