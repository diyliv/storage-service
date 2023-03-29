package timescaledb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/diyliv/store/config"
)

func ConnectToTimeScaleDb(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	timeScaleInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.TimeScaleDb.Login,
		cfg.TimeScaleDb.Password,
		cfg.TimeScaleDb.Host,
		cfg.TimeScaleDb.Port,
		cfg.TimeScaleDb.DB,
		cfg.TimeScaleDb.SSLMode)

	conn, err := pgx.Connect(ctx, timeScaleInfo)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, err
}
