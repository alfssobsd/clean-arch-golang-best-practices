package heavyprocessor

import (
	"go.uber.org/zap"
)

type HeavyProcessor struct {
	configNumber int
	logger       *zap.SugaredLogger
	pool         *heavyProcessorPool
	store        *heavyProcessorMemoryStore
}

type IHeavyProcessor interface {
	ExecuteProcessor(number int) error
	LoadNewConfigurationForProcessor(number int) error
}

func NewHeavyProcessor(logger *zap.SugaredLogger, size int) (*HeavyProcessor, error) {
	logger.Infof("Create new pool for heavyprocessor (size %d)", size)
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

func (hp *HeavyProcessor) LoadNewConfigurationForProcessor(number int) error {
	hp.logger.Infof("Set new number config to store %d", number)
	hp.store.SetNewNumberConfig(number)
	return nil
}

func (hp *HeavyProcessor) ExecuteProcessor(number int) error {
	hp.logger.Infof("Execute task with number %d", number)
	hp.logger.Infof("Pool status idle=%d/usage=%d", len(hp.pool.idle), len(hp.pool.active))
	processorItem, err := hp.pool.getProcessorItemFromPool()
	if err != nil {
		hp.logger.Errorf("No free processor idle=%d/usage=%d", len(hp.pool.idle), len(hp.pool.active))
		return err
	}
	defer func(pool *heavyProcessorPool, target *heavyProcessorPoolItem) {
		errR := pool.receiveProcessorItemToPool(target)
		if errR != nil {
			hp.logger.Error(errR)
		}
	}(hp.pool, processorItem)

	hp.logger.Infof("User processor id = %d and configNumber = %d", processorItem.GetID(), hp.store.GetNumberConfig())
	processorItem.Execute()

	return nil
}
