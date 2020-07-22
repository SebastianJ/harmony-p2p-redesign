package v1

import libp2p_pubsub "github.com/libp2p/go-libp2p-pubsub"

// Router - a catch all router/processor used in v1 of the design where multiple different message types might get propagated on a single p2p topic
type Router struct{}

func (router *Router) Process(message *libp2p_pubsub.Message) {
	// Process pubsub message here, determine how to further act upon the message using logic similar to harmony-one/harmony/node/node.go
}
