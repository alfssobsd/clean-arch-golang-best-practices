package usecases

import (
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
	"context"
	"math/rand"
)

type SystemUseCase struct {
	logger         *loggerhelper.CustomLogger
	heavyProcessor heavyprocessor.IHeavyProcessor
}

type ISystemUseCase interface {
	UpdateHeavyProcessorConfiguration(ctx context.Context) error
}

func NewSystemUseCase(logger *loggerhelper.CustomLogger, heavyProcessor heavyprocessor.IHeavyProcessor) ISystemUseCase {
	uc := SystemUseCase{
		logger:         logger,
		heavyProcessor: heavyProcessor,
	}
	return &uc
}

func (uc *SystemUseCase) UpdateHeavyProcessorConfiguration(ctx context.Context) error {
	uc.logger.InfofWithTracing(ctx, "Update Heavy Processor configuration")
	_ = uc.heavyProcessor.LoadNewConfigurationForProcessor(ctx, rand.Int())
	return nil
}
