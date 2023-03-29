package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type repository struct {
	logger        *zap.Logger
	timescaleConn *pgx.Conn
}

func NewRepository(logger *zap.Logger, timescaleConn *pgx.Conn) *repository {
	return &repository{
		logger:        logger,
		timescaleConn: timescaleConn,
	}
}

func (r *repository) Insert(ctx context.Context) {
}
