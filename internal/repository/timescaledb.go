package repository

import (
	"context"
	"fmt"

	"github.com/diyliv/store/internal/models"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type repository struct {
	logger    *zap.Logger
	timescale *pgx.Conn
}

func NewRepository(logger *zap.Logger, timescale *pgx.Conn) *repository {
	return &repository{
		logger:    logger,
		timescale: timescale,
	}
}

func (r *repository) Insert(ctx context.Context, metrics models.Response) error {
	switch t := metrics.TagValue.(type) {
	case float64:
		if _, err := r.timescale.Exec(ctx, "INSERT INTO metrics(assetId, value, quality, tstamp) VALUES($1, $2, $3, $4)",
			metrics.AssetId,
			metrics.TagValue.(float64),
			metrics.TagQuality,
			metrics.ReadAt); err != nil {
			r.logger.Error("Error while inserting data: " + err.Error())
			return err
		}
	default:
		fmt.Printf("%T\n", t)
	}
	return nil
}
