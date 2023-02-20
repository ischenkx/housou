package stream

import (
	"context"
	"io"
)

const (
	StatusOK = iota
	StatusClosed
)

type Stream interface {
	ID() string
	Send(ctx context.Context, data []byte) error
	Status() int
	// TODO: consider replacing this method with a method "Incoming() <-chan Message"
	// TODO: in Provider
	OnMessage(func(message Message))

	io.Closer
}
