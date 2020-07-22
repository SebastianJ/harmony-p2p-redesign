package p2p

import (
	"context"

	libp2p_discovery "github.com/libp2p/go-libp2p-discovery"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	libp2p_pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// Subscriber - subscription related behavior
type Subscriber interface {
	Subscribe() error
	SubscribeWithOptions(opts ...libp2p_pubsub.SubOpt) error
}

// Consumer - consumers consume messages
type Consumer interface {
	Consume() // Consume messages - depending on implementer, either consume from one given topic, or consume from all topics
}

// Processor - processors receive messages and act upon them
type Processor interface {
	Process(*libp2p_pubsub.Message)
}

// Broadcaster - broadcasters broadcast p2p messages
type Broadcaster interface {
	Broadcast(message []byte) error // Send/propagate p2p messages to a given topic
}

// Orchestrator - orchestrators configure and interact with protocol wiring etc
type Orchestrator interface {
	Rendezvous() string
	DHT() *kaddht.IpfsDHT
	Discovery() *libp2p_discovery.RoutingDiscovery
	PubSub() *libp2p_pubsub.PubSub
	Reset()                              // Reset pubsub, reset previous consumers & producers - e.g. for resharding & re-initialization of pubsub & topic consumers
	Advertise(ctx context.Context) error // spawn go routine, routinely advertise rendezvous message every x seconds
}
