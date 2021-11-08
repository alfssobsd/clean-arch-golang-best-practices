package usecases

import (
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"go.uber.org/zap"
	"math/rand"
)

type SystemUseCase struct {
	logger         *zap.SugaredLogger
	heavyProcessor heavyprocessor.IHeavyProcessor
}

type ISystemUseCase interface {
	UpdateHeavyProcessorConfiguration() error
}

func NewSystemUseCase(logger *zap.SugaredLogger, heavyProcessor heavyprocessor.IHeavyProcessor) ISystemUseCase {
	uc := SystemUseCase{
		logger:         logger,
		heavyProcessor: heavyProcessor,
	}
	return &uc
}

func (uc *SystemUseCase) UpdateHeavyProcessorConfiguration() error {
	uc.logger.Infof("Update Heavy Processor configuration")
	_ = uc.heavyProcessor.LoadNewConfigurationForProcessor(rand.Int())
	return nil
}
