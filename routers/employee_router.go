package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/resources-api/models/employee"
)

type EmployeeRouter struct {
	Storage employee.Storage
}

// CreateEmployeeHandler handles the request to create a employee
func (ep *EmployeeRouter) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	var emp employee.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee "+err.Error(), 400)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := ep.Storage.CreateEmployee(ctx, &emp)

	if err != nil {
		http.Error(w, "An error occurred when trying to enter an employee in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)

}

// GetEmployeeHandler handles the request to get a employee
func (ep *EmployeeRouter) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	employees, err := ep.Storage.GetEmployees(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employees "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)

}

// GetEmployeeByTypeHandler handles the request to get a employee by type
func (ep *EmployeeRouter) GetEmployeesByTypeHandler(w http.ResponseWriter, r *http.Request) {

	typeId := mux.Vars(r)["type_id"]

	if typeId == "" {
		http.Error(w, "An error occurred, type_id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := ep.Storage.GetEmployeeByType(ctx, typeId)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employee by type "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

}

// GetEmployeeByIdHandler handles the request to get a employee by id
func (ep *EmployeeRouter) GetEmployeesByIdHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := ep.Storage.GetEmployeeById(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get employee by id "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

}

func (ep *EmployeeRouter) UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	var e employee.Employee

	id := mux.Vars(r)["id"]

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
	employee, err := ep.Storage.UpdateEmployee(ctx, id, e)

	if err != nil {
		http.Error(w, "An error occurred when trying to update an employee in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

// DeleteEmployeeHandler handles the request to delete an employee
func (ep *EmployeeRouter) DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "An error occurred, id is required", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	employee, err := ep.Storage.DeleteEmployee(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to delete an employee in database "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

}
