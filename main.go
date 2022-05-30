package main

import (
	"fizzbuzz/handler"
	"fizzbuzz/logger"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	logger.InitLogger()
	logger.InfoLogger.Println("Fizzbuzz server initialized")
}

func main() {
	logger.InfoLogger.Println("Starting server on port 8080")

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", handler.FizzbuzzHandler)
	r.HandleFunc("/stats", handler.StatsHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
