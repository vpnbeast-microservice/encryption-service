package options

import (
	"fmt"
	"github.com/spf13/viper"
	commons "github.com/vpnbeast/golang-commons"
	"go.uber.org/zap"
	"net/http"
)

var (
	logger *zap.Logger
	opts   *EncryptionServiceOptions
)

func init() {
	logger = commons.GetLogger()
	opts = newEncryptionServiceOptions()
	err := opts.initOptions()
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

// initOptions initializes EncryptionServiceOptions while reading environment values, sets default values if not specified
func (eso *EncryptionServiceOptions) initOptions() error {
	activeProfile := commons.GetStringEnv("ACTIVE_PROFILE", "local")
	appName := commons.GetStringEnv("APP_NAME", "encryption-service")
	// TODO: below if/else logic can be implemented using library to decrease duplicate code across other projects?
	if activeProfile == "unit-test" {
		logger.Info("active profile is unit-test, reading configuration from static file")
		// TODO: better approach for that?
		viper.AddConfigPath("./../../config")
		viper.SetConfigName("unit_test")
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	} else {
		configHost := commons.GetStringEnv("CONFIG_SERVER_HOST", "localhost")
		configPort := commons.GetIntEnv("CONFIG_SERVER_PORT", 8888)
		logger.Info("loading configuration from remote server", zap.String("host", configHost),
			zap.Int("port", configPort), zap.String("appName", appName),
			zap.String("activeProfile", activeProfile))
		confAddr := fmt.Sprintf("http://%s:%d/%s-%s.yaml", configHost, configPort, appName, activeProfile)
		resp, err := http.Get(confAddr)
		if err != nil {
			return err
		}

		defer func() {
			err := resp.Body.Close()
			if err != nil {
				panic(err)
			}
		}()

		viper.SetConfigName("application")
		viper.SetConfigType("yaml")
		if err = viper.ReadConfig(resp.Body); err != nil {
			return err
		}
	}

	if err := commons.UnmarshalConfig(appName, eso); err != nil {
		return err
	}

	return nil
}
