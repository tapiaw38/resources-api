package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Func Router is a function that returns a router employees
func (ep *EmployeeRouter) EmployeeRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/create", ep.CreateEmployeeHandler).Methods("POST")
	r.HandleFunc("/all", ep.GetEmployeesHandler).Methods("GET")
	r.HandleFunc("/get_by_type/{type_id:[0-9]+}", ep.GetEmployeesByTypeHandler).Methods("GET")
	r.HandleFunc("/get_by_id/{id:[0-9]+}", ep.GetEmployeesByIdHandler).Methods("GET")
	r.HandleFunc("/update/{id:[0-9]+}", ep.UpdateEmployeeHandler).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", ep.DeleteEmployeeHandler).Methods("DELETE")

	return r
}

// Func Router is a function that returns a router workplaces
func (wp *WorkplaceRouter) WorkplaceRoutes() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/create", wp.CreateWorkplaceHandler).Methods("POST")
	r.HandleFunc("/all", wp.GetWorkplacesHandler).Methods("GET")
	r.HandleFunc("/update/{id:[0-9]+}", wp.UpdateWorkplaceHandler).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", wp.DeleteWorkplaceHandler).Methods("DELETE")

	return r
}
