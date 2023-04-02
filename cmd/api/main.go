package main

import (
	"context"

	"github.com/diyliv/store/config"
	"github.com/diyliv/store/internal/consumer"
	"github.com/diyliv/store/internal/handler"
	"github.com/diyliv/store/internal/repository"
	"github.com/diyliv/store/pkg/logger"
	"github.com/diyliv/store/pkg/storage/timescaledb"
)

func main() {
	ctx := context.Background()
	cfg := config.ReadConfig("config", "yaml", "./config")
	logger := logger.InitLogger()
	tsConn, err := timescaledb.ConnectToTimeScaleDb(ctx, cfg)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(logger, tsConn)
	kafkaConsumer := consumer.NewConsumer(cfg, logger, repo)
	defer func() {
		if err := tsConn.Close(ctx); err != nil {
			panic(err)
		}
	}()
	handler := handler.NewHandler(logger, cfg, kafkaConsumer)
	handler.Consume(ctx)
}
