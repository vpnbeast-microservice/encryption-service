package options

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"reflect"
	"strconv"
)

// unmarshalConfig creates a new *viper.Viper and unmarshalls the config into struct using *viper.Viper
func unmarshalConfig(key string, opts interface{}) error {
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	// t := reflect.TypeOf(opts)
	bindEnvs(sub, opts)
	return sub.Unmarshal(opts)
}

// bindEnvs takes *viper.Viper as argument and binds structs fields to environments variables to be able to override
// them using environment variables at the runtime
func bindEnvs(sub *viper.Viper, opts interface{}) {
	elem := reflect.ValueOf(opts).Type().Elem()
	fieldCount := elem.NumField()
	for i := 0; i < fieldCount; i++ {
		env := elem.Field(i).Tag.Get("env")
		name := elem.Field(i).Name
		_ = sub.BindEnv(name, env)
	}
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
