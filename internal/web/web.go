package web

import (
	"encryption-service/internal/options"
	"fmt"
	"github.com/gorilla/mux"
	commons "github.com/vpnbeast/golang-commons"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	logger *zap.Logger
	opts   *options.EncryptionServiceOptions
)

func init() {
	logger = commons.GetLogger()
	opts = options.GetEncryptionServiceOptions()
}

func registerHandlers(router *mux.Router) {
	encryptHandler := http.HandlerFunc(encryptHandler)
	decryptHandler := http.HandlerFunc(decryptHandler)
	checkHandler := http.HandlerFunc(checkHandler)
	pingHandler := http.HandlerFunc(pingHandler)
	router.HandleFunc("/encryption-controller/encrypt", encryptHandler).Methods("POST").
		Schemes("http").Name("encrypt")
	router.HandleFunc("/encryption-controller/decrypt", decryptHandler).Methods("POST").
		Schemes("http").Name("decrypt")
	router.HandleFunc("/encryption-controller/check", checkHandler).Methods("POST").
		Schemes("http").Name("check")
	router.HandleFunc("/health/ping", pingHandler).Methods("GET").Schemes("http").
		Name("ping")
	// router.Use(loggingMiddleware)
}

// InitServer initializes *http.Server with provided parameters
func InitServer(router *mux.Router) *http.Server {
	registerHandlers(router)
	return &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", opts.ServerPort),
		WriteTimeout: time.Duration(int32(opts.WriteTimeoutSeconds)) * time.Second,
		ReadTimeout:  time.Duration(int32(opts.ReadTimeoutSeconds)) * time.Second,
	}
}
