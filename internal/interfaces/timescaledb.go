package interfaces

import (
	"context"

	"github.com/diyliv/store/internal/models"
)

type TimeScaleDB interface {
	Insert(ctx context.Context, metrics models.Response) error
}
