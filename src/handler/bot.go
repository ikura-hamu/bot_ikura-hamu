package handler

import (
	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"go.uber.org/zap"
)

type BotHandler struct {
	br     repository.BotRepository
	cl     client.Client
	logger *zap.Logger
}

func NewBotHandler(br repository.BotRepository, cl client.Client, l *zap.Logger) *BotHandler {
	bh := &BotHandler{
		br:     br,
		cl:     cl,
		logger: l,
	}

	mc.init(bh)

	return bh
}
