package kizuna

import (
	"context"
	"kizuna/core/stream"
	"kizuna/core/subscription"
	"log"
)

type Kizuna struct {
	streamProvider stream.Provider

	streams       stream.Pool
	provider      stream.Provider
	subscriptions subscription.Pool
}

func (kizuna *Kizuna) Topic(id string) Topic {
	return Topic{
		id:     id,
		kizuna: kizuna,
	}
}

func (kizuna *Kizuna) Client(id string) Client {
	return Client{
		kizuna: kizuna,
		id:     id,
	}
}

func (kizuna *Kizuna) Run(ctx context.Context) {
	kizuna.receiveStreams(ctx)
}

func (kizuna *Kizuna) receiveStreams(ctx context.Context) {
	streams := kizuna.provider.Streams(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case newStream := <-streams:
			err := kizuna.streams.Insert(ctx, newStream)
			if err != nil {
				log.Printf("failed to create a new stream: %s\n", err)
				continue
			}
			newStream.OnMessage(func(message stream.Message) {

			})
		}
	}
}
