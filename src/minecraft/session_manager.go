package session


func NewSessionManager() (manager *SessionManager) {
	manager = new(SessionManager)
	manager.sessions = make([]*Session, 512)
	return
}


type SessionManager struct {
	sessions []*Session
}

func (manager *SessionManager) AddSession(session *Session) {
	// FIXME: grow sessions slice
	// ..
	return
}
