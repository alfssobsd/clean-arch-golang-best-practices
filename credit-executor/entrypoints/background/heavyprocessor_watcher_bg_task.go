package background

import (
	"clean-arch-golang-best-practices/credit-executor/usecases"
	"go.uber.org/zap"
	"time"
)

type HeavyProcessorWatcherBackgroundTask struct {
	logger     *zap.SugaredLogger
	systemUseCase *usecases.SystemUseCase
}

type IHeavyProcessorWatcherBackgroundTask interface {
	RunTask() error
}

func NewHeavyProcessorWatcherBackgroundTask(logger *zap.SugaredLogger, systemUseCase *usecases.SystemUseCase) *HeavyProcessorWatcherBackgroundTask {
	return &HeavyProcessorWatcherBackgroundTask{logger: logger, systemUseCase: systemUseCase}
}

func (task *HeavyProcessorWatcherBackgroundTask) RunTask() error {
	task.logger.Infof("Start HeavyProcessorWatcherBackgroundTask")
	for {
		time.Sleep(time.Second * 600)
		task.logger.Infof("Update Heavy Processor configuration")

		err := task.systemUseCase.UpdateHeavyProcessorConfiguration()
		if err != nil {
			return err
		}
	}
}

