package handshake

import session "minecraft/session"
import packets "minecraft/packets"
import packet  "minecraft/packets/handshake"
import login   "minecraft/handlers/login"
import login_p "minecraft/packets/login"

func Handler(session *session.Session, msg packets.Packet) {
	// try to cast to an Request
	_, ok := msg.(*packet.Request)
	if ok != true {
		panic("BUG!!11!!1: failed cast to Request.. something is very messed up!")
	}

	// TODO: do an database lookup

	// yup.. answer handshake now
	res := new(packet.Response)
	res.Hash = "-"

	// now add login handler
	session.SetHandler(login_p.REQ_PID, login.Handler)

	// FIXME: disable handshake?

	session.Transmit(res)
}
