package middlewares

import (
	"net/http"

	auth "github.com/tapiaw38/resources-api/routers/auth"
)

// AuthMiddleware checks if the user is logged in
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := auth.ValidateToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// AuthAdmin Middleware checks if the user is logged in and is an admin
func AuthAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, u, err := auth.ValidateToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		if !u.IsAdmin {
			http.Error(w, "Not authorized, not is admin", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
