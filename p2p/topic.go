package p2p

import (
	libp2p_pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Topic struct {
	Name         string
	Handler      *libp2p_pubsub.Topic
	Subscription *libp2p_pubsub.Subscription
}

func NewTopic(orchestrator Orchestrator, name string) (Topic, error) {
	topic := Topic{
		Name: name,
	}

	handler, err := orchestrator.PubSub().Join(name)
	if err != nil {
		return Topic{}, err
	}

	topic.Handler = handler

	return topic, nil
}

func (topic *Topic) Subscribe() error {
	if err := topic.SubscribeWithOptions(); err != nil {
		return err
	}

	return nil
}

func (topic *Topic) SubscribeWithOptions(opts ...libp2p_pubsub.SubOpt) error {
	subscription, err := topic.Handler.Subscribe(opts...)
	if err != nil {
		return err
	}

	topic.Subscription = subscription

	return nil
}
