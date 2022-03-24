package employee

import (
	"encoding/json"
	"net/http"

	employee "github.com/tapiaw38/resources-api/database/employee"
)

// GetEmployeeHandler handles the request to get a employee
func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	employees, err := employee.GetEmployees(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employees "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)

}

// GetEmployeeByTypeHandler handles the request to get a employee by type
func GetEmployeesByTypeHandler(w http.ResponseWriter, r *http.Request) {

	typeId := r.URL.Query().Get("type_id")

	if typeId == "" {
		http.Error(w, "An error occurred, type_id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := employee.GetEmployeeByType(ctx, typeId)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employee by type "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

}
