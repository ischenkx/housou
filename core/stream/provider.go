package stream

import (
	"context"
)

type Provider interface {
	Streams(ctx context.Context) <-chan Stream
}
