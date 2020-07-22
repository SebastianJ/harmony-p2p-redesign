package p2p

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type broadcasterMock struct {
	mock.Mock
}

func (mock *broadcasterMock) Broadcast(message []byte) error {
	args := mock.Called(message)
	return args.Error(0)
}

func TestBroadcasting(t *testing.T) {
	mockBroadcaster := new(broadcasterMock)

	message := []byte{0, 1, 1}

	mockBroadcaster.On("Broadcast", message).Return(nil)

	broadcaster := BroadcastSupervisor{
		Broadcasters: map[string]Broadcaster{"hmy/testnet/0.0.1/client/beacon": mockBroadcaster},
	}

	broadcaster.Broadcasters["hmy/testnet/0.0.1/client/beacon"].Broadcast(message)

	mockBroadcaster.AssertExpectations(t)
}
