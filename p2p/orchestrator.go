package p2p

import (
	"context"

	libp2p_discovery "github.com/libp2p/go-libp2p-discovery"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	libp2p_pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type StandardOrchestrator struct {
	rendezvous string
	dht        *kaddht.IpfsDHT
	discovery  *libp2p_discovery.RoutingDiscovery
	pubSub     *libp2p_pubsub.PubSub
}

func (orchestrator *StandardOrchestrator) Rendezvous() string {
	return orchestrator.rendezvous
}

func (orchestrator *StandardOrchestrator) DHT() *kaddht.IpfsDHT {
	return orchestrator.dht
}

func (orchestrator *StandardOrchestrator) Discovery() *libp2p_discovery.RoutingDiscovery {
	return orchestrator.discovery
}

func (orchestrator *StandardOrchestrator) PubSub() *libp2p_pubsub.PubSub {
	return orchestrator.pubSub
}

func (orchestrator *StandardOrchestrator) Reset() {
	orchestrator.pubSub = nil // Not full implementation, should reset current subs, topics etc propely as well
}

func (orchestrator *StandardOrchestrator) Advertise(ctx context.Context) error {
	_, err := orchestrator.discovery.Advertise(ctx, orchestrator.rendezvous)
	if err != nil {
		return err
	}

	return nil
}
