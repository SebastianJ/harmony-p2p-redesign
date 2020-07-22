package v2

import libp2p_pubsub "github.com/libp2p/go-libp2p-pubsub"

type ConsensusProcessor struct{}

func (consensusProcessor *ConsensusProcessor) Process(message *libp2p_pubsub.Message) {
	// Process pubsub message here
}
