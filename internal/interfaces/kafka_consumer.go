package interfaces

import "context"

type Consumer interface {
	Consume(ctx context.Context) error
}
