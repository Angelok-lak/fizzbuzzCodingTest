package router

import (
	"fizzbuzz/handler"
	"fizzbuzz/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func validatingMiddleware(next http.Handler) http.Handler {
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

func InitServer(port int) {

	logger.InfoLogger.Println("Starting server on port", port)

	r := mux.NewRouter()

	checkedRoutes := r.PathPrefix("/").Subrouter()

	checkedRoutes.HandleFunc("/fizzbuzz", handler.FizzbuzzHandler)
	r.HandleFunc("/stats", handler.StatsHandler)

	checkedRoutes.Use(validatingMiddleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
