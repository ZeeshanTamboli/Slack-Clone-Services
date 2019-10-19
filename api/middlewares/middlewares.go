package middlewares

import (
	"net/http"
)

// SetMiddlewareJSON : Sets content to be application/json and returns the handlerFunc
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
