package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// UserRouter is a function that returns a router for the user routes
func (ur *UserRouter) UserRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/create", ur.CreateUserHandler).Methods("POST")
	r.HandleFunc("/login", ur.LoginHandler).Methods("POST")
	r.HandleFunc("/all", ur.GetUsersHandler).Methods("GET")
	r.HandleFunc("/get_by_id/{id}", ur.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/update/{id}", ur.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", ur.DeleteUserHandler).Methods("DELETE")

	return r
}

// EmployeeRouter is a function that returns a router employees
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

// WorkplaceRouter is a function that returns a router workplaces
func (wp *WorkplaceRouter) WorkplaceRoutes() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/create", wp.CreateWorkplaceHandler).Methods("POST")
	r.HandleFunc("/all", wp.GetWorkplacesHandler).Methods("GET")
	r.HandleFunc("/update/{id:[0-9]+}", wp.UpdateWorkplaceHandler).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", wp.DeleteWorkplaceHandler).Methods("DELETE")

	return r
}

// EmployeeTypeRouter is a function that returns a router employees
func (et *EmployeeTypeRouter) EmployeeTypeRoutes() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/create", et.CreateEmployeeTypeHandler).Methods("POST")
	r.HandleFunc("/all", et.GetEmployeeTypesHandler).Methods("GET")

	return r
}

// CardRouter is a function that returns a router cards
func (c *CardRouter) CardRoutes() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/create", c.CreateCardHandler).Methods("POST")
	r.HandleFunc("/all", c.GetCardsHandler).Methods("GET")
	r.HandleFunc("/update/{id:[0-9]+}", c.UpdateCardHandler).Methods("PUT")

	return r
}
