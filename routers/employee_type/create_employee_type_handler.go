package employeetype

import (
	"encoding/json"
	"net/http"

	employeeType "github.com/tapiaw38/resources-api/database/employee_type"
	"github.com/tapiaw38/resources-api/models"
)

// CreateEmployeeTypeHandler is a handler for creating a new employee type
func CreateEmployeeTypeHandler(w http.ResponseWriter, r *http.Request) {

	var et models.EmployeeType

	err := json.NewDecoder(r.Body).Decode(&et)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee type "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employeeType, err := employeeType.CreateEmployeeType(ctx, &et)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee type in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employeeType)

}
