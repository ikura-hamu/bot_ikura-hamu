package dev

import (
	"context"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"go.uber.org/zap"

	"github.com/google/uuid"
)

var botId = conf.GetBotId()

func (dc *DevClient) JoinChannel(ctx context.Context, channelId uuid.UUID) error {
	dc.logger.Debug("join channel",
		zap.String("bot id", botId),
		zap.String("channel id", channelId.String()))
	return nil
}

func (dc *DevClient) LeaveChannel(ctx context.Context, channelId uuid.UUID) error {
	dc.logger.Debug("leave channel",
		zap.String("bot id", botId),
		zap.String("channel id", channelId.String()))
	return nil
}
