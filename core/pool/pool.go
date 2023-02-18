package pool

import (
	"context"
)

type Item interface {
	ID() string
}

type Pool[I Item, Selector any] interface {
	Insert(ctx context.Context, item I) error
	Select(ctx context.Context, selector Selector) (Set[I], error)
}

type Set[I Item] interface {
	Delete(ctx context.Context)
	Iter(ctx context.Context, iter func(I) bool)
}
