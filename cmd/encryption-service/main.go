package main

import (
	"encryption-service/pkg/metrics"
	"encryption-service/pkg/options"
	"encryption-service/pkg/web"
	"github.com/gorilla/mux"
	commons "github.com/vpnbeast/golang-commons"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	opts   *options.EncryptionServiceOptions
)

func init() {
	logger = commons.GetLogger()
	opts = options.GetEncryptionServiceOptions()
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
	server := web.InitServer(router)

	logger.Info("web server is up and running", zap.Int("serverPort", opts.ServerPort))
	panic(server.ListenAndServe())
}
