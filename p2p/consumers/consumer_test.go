package consumers

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/SebastianJ/harmony-p2p-redesign/p2p"
)

type consumerMock struct {
	mock.Mock
}

func (mock *consumerMock) Consume() {
	mock.Called()
}

func TestMessageConsumption(t *testing.T) {
	mockConsumer := new(consumerMock)
	mockConsumer.On("Consume").Return()

	consumer := ConsumerSupervisor{
		Consumers: []p2p.Consumer{
			mockConsumer,
		},
	}

	consumer.Consume()

	mockConsumer.AssertExpectations(t)
}
