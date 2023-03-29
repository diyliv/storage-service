package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"

	"github.com/diyliv/store/config"
)

func NewKafkaConn(cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialLeader(
		context.Background(),
		"tcp",
		cfg.Kafka.Brokers[0],
		cfg.Kafka.ReadFrom,
		cfg.Kafka.Partition)
}
