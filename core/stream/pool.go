package stream

import (
	"context"
	"kizuna/core/pool"
)

type Pool pool.Pool[Stream, Selector]

type Selector interface {
	Iter(ctx context.Context, iter func(string))
}

type FuncSelector struct {
	f func(ctx context.Context, iter func(string))
}

func NewFuncSelector(f func(ctx context.Context, iter func(string))) FuncSelector {
	return FuncSelector{f: f}
}

func (f FuncSelector) Iter(ctx context.Context, iter func(string)) {
	f.f(ctx, iter)
}
