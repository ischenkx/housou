package kizuna

import (
	"context"
	"github.com/google/uuid"
	"kizuna/core/stream"
	"kizuna/core/subscription"
	"log"
)

type Topic struct {
	id     string
	kizuna *Kizuna
}

func (topic Topic) Subscribe(ctx context.Context, streamID string) error {
	return topic.kizuna.subscriptions.Insert(ctx, subscription.Subscription{
		Identifier: uuid.New().String(),
		Topic:      topic.id,
		StreamID:   streamID,
	})
}

func (topic Topic) Unsubscribe(ctx context.Context, streams ...string) error {
	subs, err := topic.kizuna.subscriptions.Select(ctx, subscription.Selector{
		Streams: streams,
		Topics:  []string{topic.id},
	})
	if err != nil {
		return err
	}

	subs.Delete(ctx)
	return nil
}

func (topic Topic) Send(ctx context.Context, data []byte) error {
	subs, err := topic.kizuna.subscriptions.Select(ctx, subscription.Selector{
		Topics: []string{topic.id},
	})
	if err != nil {
		return err
	}

	streams, err := topic.kizuna.streams.Select(ctx, stream.NewFuncSelector(func(ctx context.Context, iter func(string)) {
		subs.Iter(ctx, func(sub subscription.Subscription) bool {
			iter(sub.Identifier)
			return true
		})
	}))
	if err != nil {
		return err
	}

	streams.Iter(ctx, func(s stream.Stream) bool {
		if err := s.Send(ctx, data); err != nil {
			log.Println(err)
		}
		return true
	})

	return nil
}
