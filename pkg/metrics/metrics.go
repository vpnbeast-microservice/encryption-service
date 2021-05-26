package metrics

import (
	"encryption-service/pkg/logging"
	"encryption-service/pkg/options"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	logger *zap.Logger
	opts   *options.EncryptionServiceOptions
)

func init() {
	logger = logging.GetLogger()
	opts = options.GetEncryptionServiceOptions()
}

// TODO: Generate custom metrics, check below:
// https://prometheus.io/docs/guides/go-application/
// https://www.robustperception.io/prometheus-middleware-for-gorilla-mux

// RunMetricsServer provides an endpoint, exports prometheus metrics using prometheus client golang
func RunMetricsServer(router *mux.Router) {
	metricServer := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", opts.MetricsPort),
		WriteTimeout: time.Duration(int32(opts.WriteTimeoutSeconds)) * time.Second,
		ReadTimeout:  time.Duration(int32(opts.ReadTimeoutSeconds)) * time.Second,
	}
	router.Handle("/metrics", promhttp.Handler())
	logger.Info("metric server is up and running", zap.Int("metricsPort", opts.MetricsPort))
	panic(metricServer.ListenAndServe())
}
