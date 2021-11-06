package viperhelper

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func ReadFromEnv(configPath string) {
	workingDirectory, _ := os.Getwd()
	zap.S().Infof("Try read config enviroments from %s", filepath.Join(workingDirectory, configPath))

	if len(configPath) > 0 {
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			if err == nil {
				viper.SetConfigFile(configPath)
				readInConfig()
			} else {
				zap.S().Infof("Not found env file = %s, use current enviroment variables for configuration", configPath)
				viper.AutomaticEnv()
			}
		} else {
			viper.SetConfigFile(configPath)
			readInConfig()
		}
	}
}

func readInConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			zap.S().Errorf("%s, use defaults settings", err)
		} else {
			zap.S().Fatalf("Configuration has error = %v", err)
		}
	}
}
