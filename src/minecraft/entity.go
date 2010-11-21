package entity

import "sync"

func NewEntityManager() (manager *EntityManager) {
	manager = &EntityManager{
		mapping: make(map[string]int32),
	}

	return
}


type EntityManager struct {
	mapping    map[string]int32
	lock       sync.Mutex
	currentPos int32
}

func (m *EntityManager) GetEntityID(uid string) {
	// lock this manager
	m.lock.Lock()
	defer m.lock.Unlock()

	// increment entity id by one
	m.currentPos++

	// create mapping
	m.mapping[uid] = m.currentPos

	// TODO: reverse mapping?
}
