package nats

import "context"

type Metrics interface {
	IncrementCounter(ctx context.Context, name string, labels ...string)
}
