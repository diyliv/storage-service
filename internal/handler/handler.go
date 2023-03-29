package handler

import (
	"context"

	"go.uber.org/zap"

	"github.com/diyliv/store/config"
	"github.com/diyliv/store/internal/interfaces"
)

type handler struct {
	logger   *zap.Logger
	cfg      *config.Config
	consumer interfaces.Consumer
}

func NewHandler(logger *zap.Logger, cfg *config.Config, consumer interfaces.Consumer) *handler {
	return &handler{
		logger:   logger,
		cfg:      cfg,
		consumer: consumer,
	}
}

func (h *handler) Consume(ctx context.Context) error {
	if err := h.consumer.Consume(ctx); err != nil {
		h.logger.Error("Error while consuming messages: " + err.Error())
		return err
	}
	return nil
}
