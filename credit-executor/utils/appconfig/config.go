package appconfig

import (
	"clean-arch-golang-best-practices/credit-library/viperhelper"
	"github.com/spf13/viper"
)

type IAppConfiguration interface {
	GetDatabaseConfig() DatabaseConfiguration
	GetSentryConfig() SentryConfiguration
	GetHttpServerConfig() HttpServerConfiguration
}

func NewAppConfigurationFromEnvFile(pathToEnvFile string) *AppConfiguration {
	viperhelper.ReadFromEnv(pathToEnvFile)
	appConfig := AppConfiguration{}
	appConfig.sentry = SentryConfiguration{
		Dns:         viper.GetString("SENTRY_DNS"),
		Environment: viper.GetString("SENTRY_ENVIRONMENT"),
	}
	appConfig.database = DatabaseConfiguration{
		Host:     viper.GetString("DATABASE_HOST"),
		Username: viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASSWORD"),
	}
	appConfig.httpServer = HttpServerConfiguration{
		Host: viper.GetString("SERVER_HTTP_HOST"),
		Port: viper.GetInt("SERVER_HTTP_PORT"),
	}
	return &appConfig
}

func (c *AppConfiguration) GetDatabaseConfig() DatabaseConfiguration {
	return c.database
}

func (c *AppConfiguration) GetSentryConfig() SentryConfiguration {
	return c.sentry
}

func (c *AppConfiguration) GetHttpServerConfig() HttpServerConfiguration {
	return c.httpServer
}
