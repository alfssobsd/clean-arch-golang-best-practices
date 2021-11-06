package heavyprocessor

import (
	"math/rand"
	"sync"
)

type heavyProcessorMemoryStore struct {
	numberConfig int
	mulock       *sync.Mutex
}

func NewHeavyProcessorMemoryStore() *heavyProcessorMemoryStore {
	return &heavyProcessorMemoryStore{numberConfig: rand.Int(), mulock: new(sync.Mutex)}
}

func (hpms *heavyProcessorMemoryStore) SetNewNumberConfig(newNumberConfig int) {
	hpms.mulock.Lock()
	defer hpms.mulock.Unlock()
	hpms.numberConfig = newNumberConfig
}

func (hpms *heavyProcessorMemoryStore) GetNumberConfig() int {
	hpms.mulock.Lock()
	defer hpms.mulock.Unlock()
	return hpms.numberConfig
}
