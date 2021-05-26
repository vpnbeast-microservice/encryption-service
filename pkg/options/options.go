package options

import (
	"encryption-service/pkg/logging"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var (
	logger *zap.Logger
	opts   *EncryptionServiceOptions
)

func init() {
	logger = logging.GetLogger()
	opts = newEncryptionServiceOptions()
	opts.initOptions()
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
	ServerPort          int
	MetricsPort         int
	MetricsEndpoint     string
	WriteTimeoutSeconds int
	ReadTimeoutSeconds  int
	// encryption related config
	Secret string
}

// initOptions initializes EncryptionServiceOptions while reading environment values, sets default values if not specified
func (eso *EncryptionServiceOptions) initOptions() {
	eso.ServerPort = getIntEnv("SERVER_PORT", 8085)
	eso.MetricsPort = getIntEnv("METRICS_PORT", 8086)
	eso.WriteTimeoutSeconds = getIntEnv("WRITE_TIMEOUT_SECONDS", 10)
	eso.ReadTimeoutSeconds = getIntEnv("READ_TIMEOUT_SECONDS", 10)
	eso.MetricsEndpoint = getStringEnv("METRICS_ENDPOINT", "/metrics")
	eso.Secret = getStringEnv("SECRET", "passphrasewhichneedstobe32bytes!")
}

// getStringEnv gets the specific environment variables with default value, returns default value if variable not set
func getStringEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// getIntEnv gets the specific environment variables with default value, returns default value if variable not set
func getIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return convertStringToInt(value)
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Warn("an error occurred while converting from string to int. Setting it as zero",
			zap.String("error", err.Error()))
		i = 0
	}
	return i
}
