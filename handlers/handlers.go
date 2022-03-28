package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/tapiaw38/resources-api/middlewares"
	auth "github.com/tapiaw38/resources-api/routers/auth"
	employee "github.com/tapiaw38/resources-api/routers/employee"
	employeeType "github.com/tapiaw38/resources-api/routers/employee_type"
	user "github.com/tapiaw38/resources-api/routers/user"
	workplace "github.com/tapiaw38/resources-api/routers/workplace"
)

// HandleServer handles the server request
func HandlerServer() {
	router := mux.NewRouter()

	// Initialize the routes
	users := router.PathPrefix("/api/v1/users").Subrouter()
	employees := router.PathPrefix("/api/v1/employees").Subrouter()
	workplaces := router.PathPrefix("/api/v1/workplaces").Subrouter()
	employeeTypes := router.PathPrefix("/api/v1/types").Subrouter()

	// Routes for users

	//create
	users.Path("/register").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDBMiddleware(user.CreateUserHandler))
	//get all
	users.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(middlewares.AuthAdminMiddleware(user.GetUsersHandler)))
	//get by id
	users.Path("/user/{id:[0-9]+}").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(middlewares.AuthMiddleware(user.GetUserByIdHandler)))
	//get by username
	users.Path("/profile").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(middlewares.AuthMiddleware(user.GetUserByUsernameHandler)))
	//update
	users.Path("/update/{id:[0-9]+}").Methods(
		http.MethodPut).HandlerFunc(middlewares.CheckDBMiddleware(middlewares.AuthMiddleware(user.UpdateUserHandler)))
	//delete
	users.Path("/delete/{id:[0-9]+}").Methods(
		http.MethodDelete).HandlerFunc(middlewares.CheckDBMiddleware(middlewares.AuthAdminMiddleware(user.DeleteUserHandler)))
	//login
	users.Path("/login").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDBMiddleware(auth.LoginHandler))

	// Routes for employees
	employees.Path("/create").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDBMiddleware(employee.CreateEmployeeHandler))
	employees.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(employee.GetEmployeesHandler))
	employees.Path("/update/{id:[0-9]+}").Methods(
		http.MethodPut).HandlerFunc(middlewares.CheckDBMiddleware(employee.UpdateEmployeeHandler))
	employees.Path("/get_by_type/{type_id:[0-9]+}").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(employee.GetEmployeesByTypeHandler))

	// Routes for workplaces
	workplaces.Path("/create").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDBMiddleware(workplace.CreateWorkplaceHandler))
	workplaces.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(workplace.GetWorkplacesHandler))
	workplaces.Path("/update/{id:[0-9]+}").Methods(
		http.MethodPut).HandlerFunc(middlewares.CheckDBMiddleware(workplace.UpdateWorkplaceHandler))
	workplaces.Path("/delete/{id:[0-9]+}").Methods(
		http.MethodDelete).HandlerFunc(middlewares.CheckDBMiddleware(workplace.DeleteWorkplaceHandler))

	// Routes for employee types
	employeeTypes.Path("/create").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDBMiddleware(employeeType.CreateEmployeeTypeHandler))
	employeeTypes.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDBMiddleware(employeeType.GetEmployeeTypesHandler))

	handler := cors.AllowAll().Handler(router)

	// Start the server

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: handler,
	}

	log.Println("Server started on port " + PORT)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
