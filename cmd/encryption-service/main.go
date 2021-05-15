package main

import (
	"encryption-service/pkg/config"
	"encryption-service/pkg/logging"
	"encryption-service/pkg/metrics"
	"encryption-service/pkg/web"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	serverPort, metricsPort, writeTimeoutSeconds, readTimeoutSeconds int
	logger                                                           *zap.Logger
)

func init() {
	logger = logging.GetLogger()
	serverPort = config.GetIntEnv("SERVER_PORT", 8085)
	metricsPort = config.GetIntEnv("METRICS_PORT", 8086)
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
	go metrics.RunMetricsServer(router, metricsPort, writeTimeoutSeconds, readTimeoutSeconds)
	server := web.InitServer(router, serverPort, writeTimeoutSeconds, readTimeoutSeconds)

	logger.Info("web server is up and running", zap.Int("serverPort", serverPort))
	panic(server.ListenAndServe())
}
