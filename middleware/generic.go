package middleware

import (
	"log"
	"net/http"
)

func Generic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Printf("%s - %s\n", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
