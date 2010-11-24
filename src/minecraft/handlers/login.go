package login

import session "minecraft/session"
import packets "minecraft/packets"
import packet  "minecraft/packets/login"

func Handler(session *session.Session, msg packets.Packet) {
	// try to cast to an request
	_, ok := msg.(*packet.Request)
	if ok != true {
		panic("BUG!!11!!1: failed cast to Request.. something is very messed up!")
	}

	// TODO: do something now...
}
