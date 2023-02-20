package housou

import (
	"context"
	"kizuna/core/stream"
	"kizuna/core/subscription"
)

type Client struct {
	kizuna *Housou
	id     string
}

func (client Client) Send(ctx context.Context, data []byte) error {
	streams, err := client.kizuna.streams.Select(ctx, stream.NewFuncSelector(func(ctx context.Context, iter func(string)) {
		iter(client.id)
	}))
	if err != nil {
		return err
	}

	streams.Iter(ctx, func(s stream.Stream) bool {
		_ = s.Send(ctx, data)
		return true
	})
	return nil
}

func (client Client) Subscribe(ctx context.Context, topic string) error {
	return client.kizuna.Topic(topic).Subscribe(ctx, client.id)
}

func (client Client) Unsubscribe(ctx context.Context, topics ...string) error {
	subs, err := client.kizuna.subscriptions.Select(ctx, subscription.Selector{
		Streams: []string{client.id},
		Topics:  topics,
	})
	if err != nil {
		return err
	}

	subs.Delete(ctx)
	return nil
}
