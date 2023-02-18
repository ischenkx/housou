package subscription

type Subscription struct {
	Identifier string
	Topic      string
	StreamID   string
}

func (s Subscription) ID() string {
	return s.Identifier
}
