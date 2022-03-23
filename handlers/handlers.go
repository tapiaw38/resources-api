package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/tapiaw38/resources-api/middlewares"
	employee "github.com/tapiaw38/resources-api/routers/employee"
	employeeType "github.com/tapiaw38/resources-api/routers/employee_type"
	login "github.com/tapiaw38/resources-api/routers/login"
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
	users.Path("/register").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDB(user.CreateUserHandler))
	users.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(user.GetUsersHandler))
	users.Path("/user").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(user.GetUserByIdHandler))
	users.Path("/profile").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(user.GetUserByUsernameHandler))
	users.Path("/update").Methods(
		http.MethodPut).HandlerFunc(middlewares.CheckDB(user.UpdateUserHandler))
	users.Path("/delete").Methods(
		http.MethodDelete).HandlerFunc(middlewares.CheckDB(user.DeleteUserHandler))
	users.Path("/login").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDB(login.LoginHandler))

	// Routes for employees
	employees.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(employee.GetEmployeesHandler))

	// Routes for workplaces
	workplaces.Path("/create").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDB(workplace.CreateWorkplaceHandler))
	workplaces.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(workplace.GetWorkplacesHandler))

	// Routes for employee types
	employeeTypes.Path("/create").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDB(employeeType.CreateEmployeeTypeHandler))
	employeeTypes.Path("").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(employeeType.GetEmployeeTypesHandler))

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
