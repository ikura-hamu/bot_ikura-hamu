package handler

import (
	"context"

	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	traqbot "github.com/traPtitech/traq-bot"
)

type BotHandler struct {
	br repository.BotRepository
}

func NewBotHandler(br repository.BotRepository) *BotHandler {
	return &BotHandler{
		br: br,
	}
}

func (bh *BotHandler) MessageCreatedHandler(c context.Context, payload traqbot.MessageCreatedPayload) error {
	return nil
}
