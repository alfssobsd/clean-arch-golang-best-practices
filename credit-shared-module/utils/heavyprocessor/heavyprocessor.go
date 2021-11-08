package heavyprocessor

import (
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
)

type HeavyProcessor struct {
	configNumber int
	logger       *loggerhelper.CustomLogger
	pool         *heavyProcessorPool
	store        *heavyProcessorMemoryStore
}

type IHeavyProcessor interface {
	ExecuteProcessor(ctx context.Context, number int) error
	LoadNewConfigurationForProcessor(ctx context.Context, number int) error
}

func NewHeavyProcessor(logger *loggerhelper.CustomLogger, size int) (*HeavyProcessor, error) {
	logger.InfofNoTracing("Create new pool for heavyprocessor (size %d)", size)
	pool, err := newHeavyProcessorPool(size)
	if err != nil {
		return nil, err
	}

	return &HeavyProcessor{
		logger: logger,
		pool:   pool,
		store:  NewHeavyProcessorMemoryStore(),
	}, nil
}

func (hp *HeavyProcessor) LoadNewConfigurationForProcessor(ctx context.Context, number int) error {
	hp.logger.InfofWithTracing(ctx, "Set new number config to store %d", number)
	hp.store.SetNewNumberConfig(number)
	return nil
}

func (hp *HeavyProcessor) ExecuteProcessor(ctx context.Context, number int) error {
	hp.logger.InfofWithTracing(ctx, "Execute task with number %d", number)
	hp.logger.InfofWithTracing(ctx, "Pool status idle=%d/usage=%d", len(hp.pool.idle), len(hp.pool.active))
	processorItem, err := hp.pool.getProcessorItemFromPool()
	if err != nil {
		hp.logger.ErrorfWithTracing(ctx, "No free processor idle=%d/usage=%d", len(hp.pool.idle), len(hp.pool.active))
		return err
	}
	defer func(ctx context.Context, pool *heavyProcessorPool, target *heavyProcessorPoolItem) {
		errR := pool.receiveProcessorItemToPool(target)
		if errR != nil {
			hp.logger.ErrorfWithTracing(ctx, "%s", errR)
		}
	}(ctx, hp.pool, processorItem)

	hp.logger.InfofWithTracing(ctx, "User processor id = %d and configNumber = %d", processorItem.GetID(), hp.store.GetNumberConfig())
	processorItem.Execute()

	return nil
}
