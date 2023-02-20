package housou

import (
	"context"
	"kizuna/core/event"
	"kizuna/core/stream"
	"kizuna/core/subscription"
	"log"
)

type Housou struct {
	provider      stream.Provider
	streams       stream.Pool
	subscriptions subscription.Pool
	events        event.Source
}

func (housou *Housou) Topic(id string) Topic {
	return Topic{
		id:     id,
		kizuna: housou,
	}
}

func (housou *Housou) Client(id string) Client {
	return Client{
		kizuna: housou,
		id:     id,
	}
}

func (housou *Housou) Events() event.Source {
	return housou.events
}

func (housou *Housou) Run(ctx context.Context) {
	housou.events.Emit(ctx, event.Event{Name: "run"})
	housou.receiveStreams(ctx)
}

func (housou *Housou) receiveStreams(ctx context.Context) {
	streams := housou.provider.Streams(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case newStream := <-streams:
			if err := housou.registerStream(ctx, newStream); err != nil {
				log.Printf("failed to register a new stream: %s\n", err)
			}
		}
	}
}

func (housou *Housou) registerStream(ctx context.Context, newStream stream.Stream) error {
	err := housou.streams.Insert(ctx, newStream)
	if err != nil {
		return err
	}

	housou.events.Emit(ctx, event.Event{
		Name: "new-stream",
		Data: newStream.ID(),
	})
	
	newStream.OnMessage(func(message stream.Message) {
		housou.handleMessage(ctx, message)
	})
	return nil
}

func (housou *Housou) handleMessage(ctx context.Context, message stream.Message) {
	housou.events.Emit(ctx, event.Event{
		Name: "message",
		Data: message,
	})
}
