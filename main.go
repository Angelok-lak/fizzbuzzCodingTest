package main

import (
	"fizzbuzz/logger"
	"fizzbuzz/router"
)

func init() {
	logger.InitLogger()
	logger.InfoLogger.Println("Fizzbuzz server initialized")
}

func main() {
	router.InitServer(8000)
}
