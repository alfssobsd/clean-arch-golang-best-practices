package main

import (
	"clean-arch-golang-best-practices/credit-executor/entrypoints/background"
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"context"
	"os"
	"time"
)

func MakeBackgroundTaskServer(logger *loggerhelper.CustomLogger, appConfig *appconfig.AppConfiguration, heavyProcessor heavyprocessor.IHeavyProcessor) {

	errorChain := make(chan error, 1)

	exchangeRateTask := background.NewHeavyProcessorWatcherBackgroundTask(logger, usecases.NewSystemUseCase(logger, heavyProcessor), time.Second*600)
	go func() {
		errorChain <- exchangeRateTask.RunTask(context.Background())
	}()

	exchangeRateTask2 := background.NewHeavyProcessorWatcherBackgroundTask(logger, usecases.NewSystemUseCase(logger, heavyProcessor), time.Second*500)
	go func() {
		errorChain <- exchangeRateTask2.RunTask(context.Background())
	}()

	logger.NoTracing().Info("Started all background processes")
	err := <-errorChain
	if err != nil {
		logger.NoTracing().Info("Background task is stopped, misbehavior")
		os.Exit(1)
	}
}
