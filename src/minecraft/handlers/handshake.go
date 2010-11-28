package handshake

import log "log4go"

import session "minecraft/session"
import packets "minecraft/packets"
import packet  "minecraft/packets/handshake"
import login   "minecraft/handlers/login"
import login_p "minecraft/packets/login"

func Handler(session *session.Session, msg packets.Packet) {
	// try to cast to an Request
	req, ok := msg.(*packet.Request)
	if ok != true {
		panic("BUG!!11!!1: failed cast to Request.. something is very messed up!")
	}

	// log handshake packet
	log.Info("Got handshake<%s>", req.Username)

	// answer handshake
	res := &packet.Response{
		Hash: "-",
	}

	// now add login handler
	session.SetHandler(login_p.REQ_PID, login.Handler)

	// FIXME: disable handshake?

	session.Transmit(res)
}
