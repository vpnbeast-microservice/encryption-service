package main

import (
	"encryption-service/pkg/config"
	"encryption-service/pkg/metrics"
	"encryption-service/pkg/web"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"time"
)

var port int
var writeTimeoutSeconds int
var readTimeoutSeconds int

func init() {
	port = config.GetIntEnv("PORT", 8085)
	writeTimeoutSeconds = config.GetIntEnv("WRITE_TIMEOUT_SECONDS", 10)
	readTimeoutSeconds = config.GetIntEnv("READ_TIMEOUT_SECONDS", 10)
}

func main() {
	router := mux.NewRouter()
	go metrics.RunMetricsServer(router)
	server := web.InitServer(router, fmt.Sprintf(":%d", port), time.Duration(int32(writeTimeoutSeconds)) * time.Second,
		time.Duration(int32(readTimeoutSeconds)) * time.Second)
	log.Printf("Server is listening on port %d!", port)
	log.Fatal(server.ListenAndServe())
}