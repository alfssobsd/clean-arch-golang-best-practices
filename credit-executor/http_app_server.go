package main

import (
	"clean-arch-golang-best-practices/credit-executor/dataproviders/agify_api_gateway"
	"clean-arch-golang-best-practices/credit-executor/dataproviders/redis_repository"
	httpcontroller "clean-arch-golang-best-practices/credit-executor/entrypoints/http_controllers"
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"fmt"
	"github.com/go-redis/redis/v8"
	echoPrometheus "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func MakeHttpServer(logger *loggerhelper.CustomLogger, appConfig *appconfig.AppConfiguration, heavyProcessor heavyprocessor.IHeavyProcessor, redisClient *redis.Client) {
	echoServer := echo.New()
	setupEchoMiddlewares(echoServer, logger)

	_ = httpcontroller.NewLoanCustomerHttpController(logger, echoServer,
		usecases.NewLoanCustomerUseCase(logger, heavyProcessor, agify_api_gateway.NewAgifyApiGateway(logger),
			main_db_provider.NewLoanRepository(logger, appConfig.GetDatabaseConfigForDbProvider()), redis_repository.NewSessionCacheRepository(logger, redisClient)))

	zap.S().Infof("Server start HTTP Server %s:%d", appConfig.GetHttpServerConfig().Host, appConfig.GetHttpServerConfig().Port)
	err := echoServer.Start(fmt.Sprintf(":%d", appConfig.GetHttpServerConfig().Port))
	if err != nil {
		zap.S().Fatalf("HTTP Server error %v", err)
	}
}

func setupEchoMiddlewares(echoServer *echo.Echo, logger *loggerhelper.CustomLogger) {
	echoServer.Use(loggerhelper.EchoCustomLogger(logger))
	prometheus := echoPrometheus.NewPrometheus("http_app_server", nil)
	prometheus.Use(echoServer)
}
