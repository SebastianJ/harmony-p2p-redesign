package v2

import libp2p_pubsub "github.com/libp2p/go-libp2p-pubsub"

type StakingTransactionProcessor struct{}

func (stakingTransactionProcessor *StakingTransactionProcessor) Process(message *libp2p_pubsub.Message) {
	// Process pubsub message here
}
