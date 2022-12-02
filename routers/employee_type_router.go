package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/resources-api/models/employee_type"
)

type EmployeeTypeRouter struct {
	Storage employee_type.Storage
}

// CreateEmployeeTypeHandler is a handler for creating a new employee type
func (ept *EmployeeTypeRouter) CreateEmployeeTypeHandler(w http.ResponseWriter, r *http.Request) {

	var et employee_type.EmployeeType

	err := json.NewDecoder(r.Body).Decode(&et)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee type "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employeeType, err := ept.Storage.CreateEmployeeType(ctx, &et)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee type in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employeeType)

}

// GetEmployeeTypesHandler is a handler for getting all employee types
func (ept *EmployeeTypeRouter) GetEmployeeTypesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	employeeTypes, err := ept.Storage.GetEmployeeTypes(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employee types "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeeTypes)

}
