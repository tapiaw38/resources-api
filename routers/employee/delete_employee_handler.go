package employee

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	employee "github.com/tapiaw38/resources-api/database/employee"
)

// DeleteEmployeeHandler handles the request to delete an employee
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := employee.DeleteEmployee(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete an employee in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

}
