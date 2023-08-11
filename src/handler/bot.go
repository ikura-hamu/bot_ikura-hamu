package handler

import (
	"github.com/ikura-hamu/bot_ikura-hamu/src/cache"
	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"go.uber.org/zap"
)

type BotHandler struct {
	br     repository.BotRepository
	cl     client.Client
	sc     cache.StampCache
	logger *zap.Logger
}

func NewBotHandler(br repository.BotRepository, cl client.Client, sc cache.StampCache, l *zap.Logger) *BotHandler {
	bh := &BotHandler{
		br:     br,
		cl:     cl,
		sc: sc,
		logger: l,
	}

	mc.init(bh)
	mmc.init(bh)

	return bh
}

var botUserId = conf.GetBotUserId()
