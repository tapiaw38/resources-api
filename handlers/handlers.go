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

	// Mount the employees router
	mount(router, "/employees", employees.EmployeeRoutes())
	mount(router, "/workplaces", workplaces.WorkplaceRoutes())

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

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
