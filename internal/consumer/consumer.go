package consumer

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/diyliv/store/config"
	"github.com/diyliv/store/internal/interfaces"
	"github.com/diyliv/store/internal/models"
)

type consumer struct {
	logger      *zap.Logger
	reader      *kafka.Reader
	timescaledb interfaces.TimeScaleDB
}

func NewConsumer(cfg *config.Config, logger *zap.Logger, timescaledb interfaces.TimeScaleDB) *consumer {
	return &consumer{
		logger: logger,
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:                cfg.Kafka.Brokers,
			GroupID:                cfg.Kafka.GroupID,
			Topic:                  cfg.Kafka.ReadFrom,
			MinBytes:               minBytes,
			MaxBytes:               maxBytes,
			QueueCapacity:          queueCapacity,
			HeartbeatInterval:      heartbeatInterval,
			CommitInterval:         commitInterval,
			PartitionWatchInterval: partitionWatchInterval,
			Logger:                 kafka.LoggerFunc(logger.Sugar().Debugf),
			ErrorLogger:            kafka.LoggerFunc(logger.Sugar().Errorf),
			MaxAttempts:            maxAttempts,
			Dialer: &kafka.Dialer{
				Timeout: dialTimeout,
			},
		}),
		timescaledb: timescaledb,
	}
}

func (c *consumer) Consume(ctx context.Context) error {
	var resp models.Response

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			c.logger.Error("Error while reading messages: " + err.Error())
			return err
		}

		if err := json.Unmarshal(msg.Value, &resp); err != nil {
			c.logger.Error("Error while unmarshalling: " + err.Error())
			return err
		}
	}
}
