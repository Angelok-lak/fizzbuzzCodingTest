package main

import (
	"fizzbuzz/handler"
	"fizzbuzz/logger"
	"fizzbuzz/middleware"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	logger.InitLogger()
	logger.InfoLogger.Println("Fizzbuzz server initialized")
}

func main() {
	port := 8000

	logger.InfoLogger.Println("Starting server on port", port)

	r := mux.NewRouter()

	checkedRoutes := r.PathPrefix("/").Subrouter()

	checkedRoutes.HandleFunc("/fizzbuzz", handler.FizzbuzzHandler)
	r.HandleFunc("/stats", handler.StatsHandler)

	checkedRoutes.Use(middleware.ValidatingMiddleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
