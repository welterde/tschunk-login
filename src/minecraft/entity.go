package entity

import "sync"

func NewEntityManager() (manager *EntityManager) {
	manager = &EntityManager{
		mapping:  make(map[string]uint32),
		iMapping: make([]string, 128),
	}

	return
}


type EntityManager struct {
	mapping    map[string]uint32
	iMapping   []string
	lock       sync.Mutex
	currentPos uint32
}

func (m *EntityManager) GetEntityID(uuid string) (eid uint32) {
	// lock this manager
	m.lock.Lock()
	defer m.lock.Unlock()

	// search for it..
	eid, ok := m.mapping[uuid]
	if !ok {
		// increment entity id by one
		m.currentPos++

		// create mapping
		m.mapping[uuid] = m.currentPos

		// reverse mapping
		m.fixIMapping()
		m.iMapping[m.currentPos] = uuid

		eid = m.currentPos
	}
	return
}


func (m *EntityManager) fixIMapping() {
	if m.currentPos >= uint32(cap(m.iMapping)) {
		newIM := make([]string, cap(m.iMapping)*2)
		copy(newIM, m.iMapping)
		m.iMapping = newIM
	}
}
