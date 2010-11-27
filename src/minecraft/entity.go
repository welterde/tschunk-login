package entity

import "sync"

func NewEntityManager() (manager *EntityManager) {
	manager = &EntityManager{
		mapping:  make(map[string]int32),
		iMapping: make([]int, 128),
	}

	return
}


type EntityManager struct {
	mapping    map[string]int32
	iMapping   []int
	lock       sync.Mutex
	currentPos int32
}

func (m *EntityManager) GetEntityID(uuid string) (eid int32) {
	// lock this manager
	m.lock.Lock()
	defer m.lock.Unlock()

	// search for it..
	eid, ok = m.mapping[uuid]
	if ok == nil {
		return eid
	} else {
		// increment entity id by one
		m.currentPos++

		// create mapping
		m.mapping[uuid] = m.currentPos

		// reverse mapping
		m.fixIMapping()
		m.iMapping[m.currentPos] = uuid
	}
}


func (m *EntityManager) fixIMapping() {
	if m.currentPos >= cap(m.iMapping) {
		newIM := make([]int, cap(m.iMapping)*2)
		copy(newIM, m.iMapping)
		m.iMapping = newIM
	}
}
