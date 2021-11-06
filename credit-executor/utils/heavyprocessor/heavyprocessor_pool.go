package heavyprocessor

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type heavyProcessorPoolItem struct {
	id int
}

func (hppi *heavyProcessorPoolItem) GetID() int {
	return hppi.id
}

func (hppi *heavyProcessorPoolItem) Execute() {
	time.Sleep(time.Millisecond * 300)
}

type heavyProcessorPool struct {
	idle   []*heavyProcessorPoolItem
	active []*heavyProcessorPoolItem
	size   int
	mulock *sync.Mutex
}

func newHeavyProcessorPool(size int) (*heavyProcessorPool, error) {
	pool := heavyProcessorPool{idle: []*heavyProcessorPoolItem{}, mulock: new(sync.Mutex)}
	for i := 0; i < size; i++ {
		item := heavyProcessorPoolItem{id: rand.Int()}
		pool.idle = append(pool.idle, &item)
	}
	pool.size = len(pool.idle)

	return &pool, nil
}

func (hpp *heavyProcessorPool) getProcessorItemFromPool() (*heavyProcessorPoolItem, error) {
	hpp.mulock.Lock()
	defer hpp.mulock.Unlock()
	if len(hpp.idle) == 0 {
		return nil, fmt.Errorf("No free Enforcer.Please try again later. ")
	}
	enforcer := hpp.idle[0]
	hpp.idle = hpp.idle[1:]
	hpp.active = append(hpp.active, enforcer)
	return enforcer, nil
}

func (hpp *heavyProcessorPool) receiveProcessorItemToPool(target *heavyProcessorPoolItem) error {
	hpp.mulock.Lock()
	defer hpp.mulock.Unlock()
	//search object in active
	foundTargetInActive := false
	currentActiveLength := len(hpp.active)
	for i, obj := range hpp.active {
		if obj.GetID() == target.GetID() {
			hpp.active[currentActiveLength-1], hpp.active[i] = hpp.active[i], hpp.active[currentActiveLength-1]
			hpp.active = hpp.active[:currentActiveLength-1]
			foundTargetInActive = true
			break
		}
	}

	if foundTargetInActive == false {
		return fmt.Errorf("Not found target in active pool, can't be return this enforcer. ")
	}

	//return to idle
	hpp.idle = append(hpp.idle, target)
	return nil
}
