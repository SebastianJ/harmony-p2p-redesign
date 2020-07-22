package p2p

import (
	libp2p_peer "github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

type Peer struct {
	IPAddress    string
	Port         uint16
	MultiAddress ma.Multiaddr
	Peer         libp2p_peer.ID
}
