package employeetype

import (
	"encoding/json"
	"net/http"

	employeeType "github.com/tapiaw38/resources-api/database/employee_type"
)

// GetEmployeeTypesHandler is a handler for getting all employee types
func GetEmployeeTypesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	employeeTypes, err := employeeType.GetEmployeeTypes(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employee types "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeeTypes)

}
