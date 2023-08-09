package handler

import (
	"context"

	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
	"github.com/ikura-hamu/bot_ikura-hamu/src/payload"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"go.uber.org/zap"
)

type BotHandler struct {
	br     repository.BotRepository
	cl     client.Client
	logger *zap.Logger
}

func NewBotHandler(br repository.BotRepository, cl client.Client, l *zap.Logger) *BotHandler {
	return &BotHandler{
		br:     br,
		cl:     cl,
		logger: l,
	}
}

func (bh *BotHandler) MessageCreatedHandler(c context.Context, payload payload.EventMessagePayload) error {
	return bh.cl.SendMessage(c, payload.ChannelID, payload.PlainText, false)
}
