package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// mapstructure
// https://github.com/mitchellh/mapstructure
// It helps to decode data from a map format into a Go struct
type AppConfig struct {
	DbURL        string `mapstructure:"DB_URL"`
	DbDriver     string `mapstructure:"DB_DRIVER"`
	ServeAddress string `mapstructure:"SERVE_ADDRESS"`
	ServePort    string `mapstructure:"SERVE_PORT"`
}

func LoadAppConfig(path string) (AppConfig, error) {
	if path == "" {
		return AppConfig{}, fmt.Errorf("config path is empty")
	}

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return AppConfig{}, fmt.Errorf("config file not found: %w", err)
		}
		return AppConfig{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return config, nil

	// https://github.com/spf13/viper
	// viper.SetConfigName("config")         // name of config file (without extension)
	// viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	// viper.AddConfigPath(".")              // optionally look for config in the working directory
	// err := viper.ReadInConfig()           // Find and read the config file
	// if err != nil {                       // Handle errors reading the config file
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }
}
