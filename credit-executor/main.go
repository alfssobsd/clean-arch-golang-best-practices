package main

import (
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-executor/utils/heavyprocessor"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
)

func main() {
	logger := loggerhelper.MakeNewZapProductionLogger()
	appConfig := appconfig.NewAppConfigurationFromEnvFile(".env")
	heavyProcessor, _ := heavyprocessor.NewHeavyProcessor(logger.Sugar(), 10)

	//TODO: обработчик который стопит весь процесс
	go func() {
		MakeBackgroundTaskServer(logger.Sugar(), appConfig, heavyProcessor)
	}()
	MakeHttpServer(logger.Sugar(), appConfig, heavyProcessor)
}
