package background

import (
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
	"math/rand"
	"time"
)

type HeavyProcessorWatcherBackgroundTask struct {
	watcherId     int
	logger        *loggerhelper.CustomLogger
	systemUseCase usecases.ISystemUseCase
	waitTimeout   time.Duration
}

type IHeavyProcessorWatcherBackgroundTask interface {
	RunTask(ctx context.Context) error
}

func NewHeavyProcessorWatcherBackgroundTask(logger *loggerhelper.CustomLogger, systemUseCase usecases.ISystemUseCase, waitTimeout time.Duration) IHeavyProcessorWatcherBackgroundTask {
	return &HeavyProcessorWatcherBackgroundTask{watcherId: rand.Int(), logger: logger, systemUseCase: systemUseCase, waitTimeout: waitTimeout}
}

func (task *HeavyProcessorWatcherBackgroundTask) RunTask(ctx context.Context) error {
	task.logger.InfofWithTracing(ctx, "Watcher = %d, Start HeavyProcessorWatcherBackgroundTask with timeout = %s", task.watcherId, task.waitTimeout.String())
	for {
		time.Sleep(task.waitTimeout)
		task.logger.InfofWithTracing(ctx, "Watcher = %d, Update Heavy Processor configuration", task.watcherId)

		err := task.systemUseCase.UpdateHeavyProcessorConfiguration(ctx)
		if err != nil {
			return err
		}
	}
}
