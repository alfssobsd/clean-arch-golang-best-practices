package main

import (
	"clean-arch-golang-best-practices/credit-executor/dataproviders/agify_api_gateway"
	httpcontroller "clean-arch-golang-best-practices/credit-executor/entrypoints/http"
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/echohelper"
	"clean-arch-golang-best-practices/credit-shared-module/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"fmt"
	echoPrometheus "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func MakeHttpServer(logger *zap.Logger, appConfig *appconfig.AppConfiguration, heavyProcessor *heavyprocessor.HeavyProcessor) {
	echoServer := echo.New()
	setupEchoMiddlewares(echoServer, logger)

	_ = httpcontroller.NewLoanCustomerHttpController(logger.Sugar(), echoServer,
		usecases.NewLoanCustomerUseCase(logger.Sugar(), heavyProcessor, agify_api_gateway.NewAgifyApiGateway(logger.Sugar()),
			main_db_provider.NewLoanRepository(logger.Sugar(), appConfig.GetDatabaseConfigForDbProvider())))

	zap.S().Infof("Server start HTTP Server %s:%d", appConfig.GetHttpServerConfig().Host, appConfig.GetHttpServerConfig().Port)
	err := echoServer.Start(fmt.Sprintf(":%d", appConfig.GetHttpServerConfig().Port))
	if err != nil {
		zap.S().Fatalf("HTTP Server error %v", err)
	}
}

func setupEchoMiddlewares(echoServer *echo.Echo, logger *zap.Logger) {
	echoServer.Use(echohelper.EchoZapLogger(logger))
	prometheus := echoPrometheus.NewPrometheus("http_app_server", nil)
	prometheus.Use(echoServer)
}
