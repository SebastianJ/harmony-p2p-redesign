package p2p

import "context"

type TopicBroadcaster struct {
	Topic   Topic
	Context context.Context
}

type BroadcastSupervisor struct {
	Broadcasters map[string]Broadcaster
}

func NewBroadcastSupervisor(ctx context.Context, orchestrator Orchestrator, topics []string) (BroadcastSupervisor, error) {
	supervisor := BroadcastSupervisor{
		Broadcasters: make(map[string]Broadcaster),
	}

	for _, topicName := range topics {
		topic, err := NewTopic(orchestrator, topicName)
		if err != nil {
			return BroadcastSupervisor{}, err
		}

		var broadcaster Broadcaster
		broadcaster = &TopicBroadcaster{
			Topic:   topic,
			Context: ctx,
		}

		supervisor.Broadcasters[topicName] = broadcaster
	}

	return supervisor, nil
}

func (topicBroadcaster *TopicBroadcaster) Broadcast(message []byte) error {
	if err := topicBroadcaster.Topic.Handler.Publish(topicBroadcaster.Context, message); err != nil {
		return err
	}

	return nil
}
