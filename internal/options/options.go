package options

import (
	commons "github.com/vpnbeast/golang-commons"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	opts   *EncryptionServiceOptions
)

func init() {
	logger = commons.GetLogger()
	opts = newEncryptionServiceOptions()
	err := commons.InitOptions(opts, "encryption-service")
	if err != nil {
		logger.Fatal("fatal error occured while initializing options", zap.Error(err))
	}
}

// GetEncryptionServiceOptions returns the initialized EncryptionServiceOptions
func GetEncryptionServiceOptions() *EncryptionServiceOptions {
	return opts
}

// newEncryptionServiceOptions creates an AuthServiceOptions struct with zero values
func newEncryptionServiceOptions() *EncryptionServiceOptions {
	return &EncryptionServiceOptions{}
}

type EncryptionServiceOptions struct {
	// web server related config
	ServerPort          int    `env:"SERVER_PORT"`
	MetricsPort         int    `env:"METRICS_PORT"`
	MetricsEndpoint     string `env:"METRICS_ENDPOINT"`
	WriteTimeoutSeconds int    `env:"WRITE_TIMEOUT_SECONDS"`
	ReadTimeoutSeconds  int    `env:"READ_TIMEOUT_SECONDS"`
	// encryption related config
	Secret string `env:"SECRET"`
}
