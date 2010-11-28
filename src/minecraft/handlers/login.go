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
		// TODO: add more handlers
	} else if req.ProtocolVersion == 7 {
		// manic digger client
	} else {
		// neither.... hmmmm... what to do now?
		// for now lets just continue and pretend it didn't happen ok?
	}

	// check username... range is 1..42
	if len(req.Username) < 1 || len(req.Username) > 42 {
		// invalid username
		// kick user
		session.Kick("Invalid username! Name too short/long!")
		// TODO: do something more useful here?
	}

	// TODO: make sure that the username doesn't contain invalid characters

	// ID of the player is /world/players/$username
	id := "/world/players/" + req.Username

	// TODO: perform authentication

	// get the EID of this ID
	eid := session.EntityManager.GetEntityID(id)

	// send response
	// TODO: what does MapSeed do?
	// TODO: support multiple dimensions
	// FIXME: is it uint32 or just int32? do we care?
	resp := &packet.Response{
		EID:       int32(eid),
		MapSeed:   0,
		Dimension: 0,
	}

	// now transmit the response
	session.Transmit(resp)
}
