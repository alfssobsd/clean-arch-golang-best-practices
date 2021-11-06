package background

import (
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type HeavyProcessorWatcherBackgroundTask struct {
	watcherId     int
	logger        *zap.SugaredLogger
	systemUseCase *usecases.SystemUseCase
	waitTimeout   time.Duration
}

type IHeavyProcessorWatcherBackgroundTask interface {
	RunTask() error
}

func NewHeavyProcessorWatcherBackgroundTask(logger *zap.SugaredLogger, systemUseCase *usecases.SystemUseCase, waitTimeout time.Duration) *HeavyProcessorWatcherBackgroundTask {
	return &HeavyProcessorWatcherBackgroundTask{watcherId: rand.Int(), logger: logger, systemUseCase: systemUseCase, waitTimeout: waitTimeout}
}

func (task *HeavyProcessorWatcherBackgroundTask) RunTask() error {
	task.logger.Infof("Watcher = %d, Start HeavyProcessorWatcherBackgroundTask with timeout = %s", task.watcherId, task.waitTimeout.String())
	for {
		time.Sleep(task.waitTimeout)
		task.logger.Infof("Watcher = %d, Update Heavy Processor configuration", task.watcherId)

		err := task.systemUseCase.UpdateHeavyProcessorConfiguration()
		if err != nil {
			return err
		}
	}
}
