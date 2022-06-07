package router

import (
	"fizzbuzz/handler"
	"fizzbuzz/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer(port int) {

	logger.InfoLogger.Println("Starting server on port", port)

	r := mux.NewRouter()

	checkedRoutes := r.PathPrefix("/").Subrouter()

	checkedRoutes.HandleFunc("/fizzbuzz", handler.FizzbuzzHandler)
	r.HandleFunc("/stats", handler.StatsHandler)

	checkedRoutes.Use(validatingMiddleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
