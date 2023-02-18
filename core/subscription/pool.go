package subscription

import "kizuna/core/pool"

type Pool pool.Pool[Subscription, Selector]

type Selector struct {
	Streams []string
	Topics  []string
}
