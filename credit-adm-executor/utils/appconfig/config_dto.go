package appconfig

type AppConfiguration struct {
	database   DatabaseConfiguration
	sentry     SentryConfiguration
	httpServer HttpServerConfiguration
}

type DatabaseConfiguration struct {
	Username string
	Password string
	Host     string
}

type SentryConfiguration struct {
	Dns         string
	Environment string
}

type HttpServerConfiguration struct {
	Host string
	Port int
}
