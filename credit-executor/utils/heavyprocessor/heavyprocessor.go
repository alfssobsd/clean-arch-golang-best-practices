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
	_ = hp.LoadNewConfigurationForProcessor(number)
	return nil
}

func (hp *HeavyProcessor) ExecuteProcessor(number int) error {
	hp.logger.Infof("Execute task with number %d", number)
	processorItem, err := hp.pool.getProcessorItemFromPool()
	if err != nil {
		return err
	}
	defer hp.pool.receiveProcessorItemToPool(processorItem)
	hp.logger.Infof("User processor id = %d and configNumber = %d", processorItem.GetID(), hp.store.GetNumberConfig())
	processorItem.Execute()

	return nil
}
