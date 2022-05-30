package handler

import (
	"fizzbuzz/fizzbuzz"
	"fizzbuzz/logger"
	"fizzbuzz/stats"

	"encoding/json"
	"fmt"
	"net/http"
)

func FizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	var s fizzbuzz.FizzbuzzStruct

	logger.InfoLogger.Println("/fizzbuzz request received with parameters:", r.Body)

	// decode json request parameters
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		logger.ErrorLogger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check extracted struct params
	fizzbuzzParams, err := fizzbuzz.CheckParams(s)
	if err != nil {
		if err != nil {
			logger.ErrorLogger.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// run fizzbuzz algorithm
	a := fizzbuzz.Fizzbuzz(fizzbuzzParams)

	// add result to stats
	stats.AddRecord(fizzbuzzParams)

	// return result
	logger.InfoLogger.Printf(fmt.Sprintf("SUCCESS: %v", a))
	response, err := json.Marshal(a)
	if err != nil {
		logger.ErrorLogger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
