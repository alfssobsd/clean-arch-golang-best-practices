package appconfig

type AppConfiguration struct {
	database   DatabaseConfiguration
	sentry     SentryConfiguration
	httpServer HttpServerConfiguration
	redis      RedisConfiguration
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

type RedisConfiguration struct {
	Host     string
	Port     int
	Password string
	Database int
	PoolSize int
}

type HttpServerConfiguration struct {
	Host string
	Port int
}
