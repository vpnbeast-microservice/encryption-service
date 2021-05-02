package metrics

import (
	"encryption-service/pkg/logging"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	logger *zap.Logger
)

func init() {
	logger = logging.GetLogger()
}

// TODO: Generate custom metrics, check below:
// https://prometheus.io/docs/guides/go-application/
// https://www.robustperception.io/prometheus-middleware-for-gorilla-mux

func RunMetricsServer(router *mux.Router, metricsPort, writeTimeoutSeconds, readTimeoutSeconds int) {
	metricServer := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(":%d", metricsPort),
		WriteTimeout: time.Duration(int32(writeTimeoutSeconds)) * time.Second,
		ReadTimeout:  time.Duration(int32(readTimeoutSeconds)) * time.Second,
	}
	router.Handle("/metrics", promhttp.Handler())
	logger.Info("metric server is up and running", zap.Int("metricsPort", metricsPort))
	panic(metricServer.ListenAndServe())
}