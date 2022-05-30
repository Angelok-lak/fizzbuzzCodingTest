package handler

import (
	"fizzbuzz/logger"
	"fizzbuzz/stats"

	"encoding/json"
	"net/http"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("/stats request received")

	// return stats
	topRequest, err := stats.GetTopRequest()
	if err != nil {
		logger.ErrorLogger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(topRequest)
	if err != nil {
		logger.ErrorLogger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
