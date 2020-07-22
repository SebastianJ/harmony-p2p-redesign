package consumers

import (
	"github.com/SebastianJ/harmony-p2p-redesign/p2p"
)

type ConsumerSupervisor struct {
	Consumers []p2p.Consumer
}

func (consumer *ConsumerSupervisor) Consume() {
	for _, topicConsumer := range consumer.Consumers {
		topicConsumer.Consume()
	}
}
