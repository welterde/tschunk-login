package handshake

import . "minecraft/packets/base"

type HandshakeRequest struct {
	username string
}

type HandshakeResponse struct {
	hash string
}

