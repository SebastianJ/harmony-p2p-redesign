package consumers

import (
	"github.com/SebastianJ/harmony-p2p-redesign/p2p"
)

type Supervisor struct {
	Consumers []p2p.Consumer
}

func (supervisor *Supervisor) Consume() {
	for _, consumer := range supervisor.Consumers {
		consumer.Consume()
	}
}
