package login

import session "minecraft/session"
import packets "minecraft/packets"
import packet  "minecraft/packets/login"

func Handler(session *session.Session, msg packets.Packet) {
	// try to cast to an request
	req, ok := msg.(*packet.Request)
	if ok != true {
		panic("BUG!!11!!1: failed cast to Request.. something is very messed up!")
	}

	// check protocol version.. it should either be 5(minecraft client) or 7(manic digger)
	if req.ProtocolVersion == 5 {
		// minecraft client
	} else if req.ProtocolVersion == 7 {
		// manic digger client
	} else {
		// neither.... hmmmm... what to do now?
		// for now lets just continue and pretend it didn't happen ok?
	}

	// TODO: get the entity-id of the user
}
