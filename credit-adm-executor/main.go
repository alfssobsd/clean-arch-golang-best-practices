package main

import (
	"clean-arch-golang-best-practices/credit-adm-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
)

func main() {
	logger := loggerhelper.NewCustomLogger()
	appConfig := appconfig.NewAppConfigurationFromEnvFile(".env")

	logger.NoTracing().Info("Run credit-adm-executor")
	MakeHttpServer(logger, appConfig)
}
