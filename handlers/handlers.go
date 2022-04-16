package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/resources-api/routers"
	"github.com/tapiaw38/resources-api/storage"
)

// HandleServer handles the server request
func HandlerServer() {
	router := mux.NewRouter()

	employees := &routers.EmployeeRouter{
		Storage: &storage.EmployeeStorage{
			Data: storage.NewConnection(),
		},
	}

	workplaces := &routers.WorkplaceRouter{
		Storage: &storage.WorkplaceStorage{
			Data: storage.NewConnection(),
		},
	}

	employeeTypes := &routers.EmployeeTypeRouter{
		Storage: &storage.EmployeeTypeStorage{
			Data: storage.NewConnection(),
		},
	}

	cards := &routers.CardRouter{
		Storage: &storage.CardStorage{
			Data: storage.NewConnection(),
		},
	}

	users := &routers.UserRouter{
		Storage: &storage.UserStorage{
			Data: storage.NewConnection(),
		},
	}

	// Mount the employees router

	mount(router, "/users", users.UserRoutes())
	mount(router, "/employees", employees.EmployeeRoutes())
	mount(router, "/workplaces", workplaces.WorkplaceRoutes())
	mount(router, "/types", employeeTypes.EmployeeTypeRoutes())
	mount(router, "/cards", cards.CardRoutes())

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

// mount is a helper function to mount a router to a path
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
