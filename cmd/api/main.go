package main

import (
	"context"

	"github.com/diyliv/store/config"
	"github.com/diyliv/store/pkg/kafka"
	"github.com/diyliv/store/pkg/storage/timescaledb"
)

func main() {
	ctx := context.Background()
	cfg := config.ReadConfig("config", "yaml", "./config")
	tsConn, err := timescaledb.ConnectToTimeScaleDb(ctx, cfg)
	if err != nil {
		panic(err)
	}
	kafkaConn, err := kafka.NewKafkaConn(cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tsConn.Close(ctx); err != nil {
			panic(err)
		}
		if err := kafkaConn.Close(); err != nil {
			panic(err)
		}
	}()
}
