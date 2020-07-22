package p2p

import (
	"github.com/harmony-one/bls/ffi/go/bls"
	libp2p_host "github.com/libp2p/go-libp2p-core/host"
	libp2p_network "github.com/libp2p/go-libp2p-core/network"
)

type Host struct {
	SelfPeer           Peer
	Peers              []Peer
	Host               libp2p_host.Host
	ConsensusPublicKey *bls.PublicKey
}

func (host *Host) Connectivity() (int, int, int) {
	connected, not := 0, 0
	peers := host.Host.Peerstore().Peers()
	for _, peer := range peers {
		result := host.Host.Network().Connectedness(peer)
		if result == libp2p_network.Connected {
			connected++
		} else if result == libp2p_network.NotConnected {
			not++
		}
	}
	return len(peers), connected, not
}
