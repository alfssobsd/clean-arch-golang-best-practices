package main

import (
	"clean-arch-golang-best-practices/credit-executor/entrypoints/background"
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"go.uber.org/zap"
	"os"
	"time"
)

func MakeBackgroundTaskServer(logger *zap.SugaredLogger, appConfig *appconfig.AppConfiguration, heavyProcessor *heavyprocessor.HeavyProcessor) {

	errorChain := make(chan error, 1)

	exchangeRateTask := background.NewHeavyProcessorWatcherBackgroundTask(logger, usecases.NewSystemUseCase(logger, heavyProcessor), time.Second*600)
	go func() {
		errorChain <- exchangeRateTask.RunTask()
	}()

	exchangeRateTask2 := background.NewHeavyProcessorWatcherBackgroundTask(logger, usecases.NewSystemUseCase(logger, heavyProcessor), time.Second*500)
	go func() {
		errorChain <- exchangeRateTask2.RunTask()
	}()

	logger.Infof("Started all background processes")
	err := <-errorChain
	if err != nil {
		logger.Error("Background task is stopped, misbehavior")
		os.Exit(1)
	}
}
