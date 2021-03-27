package main

import (
	"encryption-service/pkg/config"
	"encryption-service/pkg/logging"
	"encryption-service/pkg/metrics"
	"encryption-service/pkg/web"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"time"
)

var (
	port, writeTimeoutSeconds, readTimeoutSeconds int
	logger *zap.Logger
)

func init() {
	logger = logging.GetLogger()
	port = config.GetIntEnv("PORT", 8085)
	writeTimeoutSeconds = config.GetIntEnv("WRITE_TIMEOUT_SECONDS", 10)
	readTimeoutSeconds = config.GetIntEnv("READ_TIMEOUT_SECONDS", 10)
}

func main() {
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	router := mux.NewRouter()
	go metrics.RunMetricsServer(router)
	server := web.InitServer(router, fmt.Sprintf(":%d", port), time.Duration(int32(writeTimeoutSeconds)) * time.Second,
		time.Duration(int32(readTimeoutSeconds)) * time.Second)

	logger.Info("web server is up and running", zap.Int("port", port))
	panic(server.ListenAndServe())
}