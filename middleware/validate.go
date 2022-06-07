package middleware

import (
	"log"
	"net/http"
)

func ValidatingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		if r.Header.Get("Content-Type") != "application/json" {
			log.Println("Invalid content type")
			http.Error(w, "Invalid content type", http.StatusBadRequest)
			return
		}
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
