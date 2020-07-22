package consumers

import (
	"context"

	"golang.org/x/sync/semaphore"

	"github.com/SebastianJ/harmony-p2p-redesign/p2p"
)

type TopicConsumer struct {
	Topic     p2p.Topic
	Processor p2p.Processor
	Context   context.Context
	Semaphore *semaphore.Weighted
}

func NewTopicConsumer(ctx context.Context, topic p2p.Topic, processor p2p.Processor, semaphoreWeight int64) TopicConsumer {
	return TopicConsumer{
		Topic:     topic,
		Processor: processor,
		Context:   ctx,
		Semaphore: semaphore.NewWeighted(semaphoreWeight),
	}
}

func (topicConsumer *TopicConsumer) Consume() {
	go func() {
		for {
			if topicConsumer.Semaphore.TryAcquire(1) {
				defer topicConsumer.Semaphore.Release(1)

				// TODO: Dummy impl, should hook into individual consumers etc - add more examples later
				message, err := topicConsumer.Topic.Subscription.Next(topicConsumer.Context)
				if err != nil {
					continue
				}

				topicConsumer.Processor.Process(message)
			}
		}
	}()
}
