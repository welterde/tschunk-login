package bans

import "net"


func NewBanManager() (bans *BanManager) {
	bans = &BanManager{}
	return
}


type BanManager struct{}

func (bans *BanManager) CheckAddress(addr net.Addr) (ok bool, msg string) {
	// TODO: check if the user is banned
	ok = true
	return
}
