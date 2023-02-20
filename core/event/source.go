package event

import (
	"context"
	"io"
)

type Source interface {
	Emit(ctx context.Context, event Event)
	On(name string, handler func(ctx context.Context, event Event)) io.Closer
}
