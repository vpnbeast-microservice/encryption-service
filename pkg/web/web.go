package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

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

func InitServer(router *mux.Router, serverPort, writeTimeoutSeconds, readTimeoutSeconds int) *http.Server {
	registerHandlers(router)
	return &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(":%d", serverPort),
		WriteTimeout: time.Duration(int32(writeTimeoutSeconds)) * time.Second,
		ReadTimeout:  time.Duration(int32(readTimeoutSeconds)) * time.Second,
	}
}