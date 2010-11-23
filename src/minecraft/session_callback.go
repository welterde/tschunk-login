package session

type Daemon interface {
	SessionManager() *SessionManager
}
