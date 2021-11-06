package main

import (
	"clean-arch-golang-best-practices/credit-adm-executor/entrypoints/http_controllers"
	"clean-arch-golang-best-practices/credit-adm-executor/usecases"
	"clean-arch-golang-best-practices/credit-adm-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/echohelper"
	"clean-arch-golang-best-practices/credit-shared-module/dataproviders/main_db_provider"
	"clean-arch-golang-best-practices/credit-shared-module/domain"
	"fmt"
	echoPrometheus "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func MakeHttpServer(logger *zap.Logger, appConfig *appconfig.AppConfiguration) {
	echoServer := echo.New()
	setupEchoMiddlewares(echoServer, logger)

	_ = http_controllers.NewLoanCustomerAdmHttpController(logger.Sugar(), echoServer,
		usecases.NewLoanCustomerAdmUseCase(
			logger.Sugar(),
			domain.NewCreditRatingDomain(logger.Sugar()),
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
