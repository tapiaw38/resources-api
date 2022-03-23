package employee

import (
	"encoding/json"
	"net/http"

	employee "github.com/tapiaw38/resources-api/database/employee"
)

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
