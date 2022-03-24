package employee

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/resources-api/models"

	employee "github.com/tapiaw38/resources-api/database/employee"
)

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	var emp models.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := employee.CreateEmployee(ctx, &emp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)

}
