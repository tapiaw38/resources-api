package middlewares

import (
	"net/http"

	auth "github.com/tapiaw38/resources-api/routers/auth"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := auth.BasicAuth(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
