package main

import (
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-executor/utils/helpers"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
)

func main() {
	logger := loggerhelper.NewCustomLogger()
	appConfig := appconfig.NewAppConfigurationFromEnvFile(".env")
	heavyProcessor, _ := heavyprocessor.NewHeavyProcessor(logger, 10)
	redisClient := helpers.NewRedisClient(logger, appConfig.GetRedisConfig())

	logger.SugarNoTracing().Infof("Run credit-executor")
	go func() {
		MakeBackgroundTaskServer(logger, appConfig, heavyProcessor)
	}()
	MakeHttpServer(logger, appConfig, heavyProcessor, redisClient)
}
