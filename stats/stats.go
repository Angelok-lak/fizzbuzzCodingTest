package stats

import (
	"encoding/json"
	"errors"
	"fizzbuzz/fizzbuzz"
	"fizzbuzz/logger"
)

var (
	statsLogs map[string]int
)

type StatsResponse struct {
	Request fizzbuzz.FizzbuzzStruct `json:"request"`
	Hits    int                     `json:"hits"`
}

func InitStats() {
	// init stats array
	statsLogs = make(map[string]int)
}

func AddRecord(v interface{}) {
	if statsLogs == nil {
		InitStats()
	}

	// encode json and use it as key to improve search performance
	key, err := json.Marshal(v)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	// add record to stats
	record, ok := statsLogs[string(key)]
	if ok {
		statsLogs[string(key)] = record + 1
	} else {
		statsLogs[string(key)] = 1
	}
}

func GetTopRequest() (interface{}, error) {
	// get top request
	topValue := 1
	topKey := ""

	for k, v := range statsLogs {
		if v > topValue {
			topValue = v
			topKey = k
		}
	}

	// decode json and create response
	buffer := &StatsResponse{}
	buffer.Hits = topValue
	response := json.Unmarshal([]byte(topKey), &buffer.Request)
	if response == nil {
		return nil, errors.New("error unmarshalling json")
	} else {
		return buffer, nil
	}
}
