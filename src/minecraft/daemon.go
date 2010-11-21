package daemon

import session "minecraft/session"


func NewDaemon() (daemon *Daemon) {
	daemon = &Daemon{
		sessionManager: &session.SessionManager{},
	}
	return
}

type Daemon struct {
	sessionManager *session.SessionManager
}

func (daemon *Daemon) SessionManager() (manager *session.SessionManager) {
	return daemon.sessionManager
}

// Compile-time assertion
var _ session.Daemon = (*Daemon)(nil)
