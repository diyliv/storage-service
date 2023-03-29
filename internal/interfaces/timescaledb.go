package interfaces

import "context"

type TimeScaleDB interface {
	Insert(ctx context.Context)
}
