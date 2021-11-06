package main

import (
	"clean-arch-golang-best-practices/credit-executor/entrypoints/background"
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-executor/utils/heavyprocessor"
	"go.uber.org/zap"
)

func MakeBackgroundTaskServer(logger *zap.SugaredLogger, appConfig *appconfig.AppConfiguration, heavyProcessor *heavyprocessor.HeavyProcessor) {

	exchangeRateTask := background.NewHeavyProcessorWatcherBackgroundTask(logger, usecases.NewSystemUseCase(logger, heavyProcessor))
	go func() {
		_ = exchangeRateTask.RunTask()
	}()
}
