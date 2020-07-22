package p2p

const (
	// ProtocolVersion determines how e.g. messages should be routed
	ProtocolVersion = 1
)

// Messenger - the main entry point for p2p
type Messenger struct {
	Orchestrator Orchestrator
	Consumer     Consumer
	Broadcaster  Broadcaster
}
