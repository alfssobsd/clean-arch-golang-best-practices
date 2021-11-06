package main

import (
	"clean-arch-golang-best-practices/credit-executor/dataproviders/agify_api_gateway"
	"clean-arch-golang-best-practices/credit-executor/dataproviders/main_db_provider"
	httpcontroller "clean-arch-golang-best-practices/credit-executor/entrypoints/http"
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-executor/utils/heavyprocessor"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func MakeHttpServer(logger *zap.SugaredLogger, appConfig *appconfig.AppConfiguration, heavyProcessor *heavyprocessor.HeavyProcessor) {
	echoServer := echo.New()

	_ = httpcontroller.NewLoanCustomerHttpController(logger, echoServer,
		usecases.NewLoanCustomerUseCase(logger, heavyProcessor, agify_api_gateway.NewAgifyApiGateway(logger),
			main_db_provider.NewLoanRepository(logger, appConfig.GetDatabaseConfig())))

	zap.S().Infof("Server start HTTP Server %s:%d", appConfig.GetHttpServerConfig().Host, appConfig.GetHttpServerConfig().Port)
	err := echoServer.Start(fmt.Sprintf(":%d", appConfig.GetHttpServerConfig().Port))
	if err != nil {
		zap.S().Fatalf("HTTP Server error %v", err)
	}
}
