package dev

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const MESSAGE_MSG = "message"

func (dc *DevClient) SendMessage(ctx context.Context, channelId uuid.UUID, message string, embed bool) (uuid.UUID, error) {
	mesId := uuid.New()
	dc.logger.Debug(MESSAGE_MSG,
		zap.String("channel id", channelId.String()),
		zap.String("message id", mesId.String()),
		zap.String("message", message),
		zap.Bool("embed", embed))
	return mesId, nil
}
