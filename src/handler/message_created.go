package handler

import (
	"context"
	"regexp"
	"sync"

	"github.com/ikura-hamu/bot_ikura-hamu/src/payload"
	"go.uber.org/zap"
)

type messageCreated struct {
	mc map[string]func(ctx context.Context, payload payload.EventMessagePayload) error
	m  sync.Mutex
}

var mc messageCreated

func (bh *BotHandler) MessageCreatedHandler(ctx context.Context, payload payload.EventMessagePayload) error {
	mc.m.Lock()
	defer mc.m.Unlock()

	bh.logger.Debug("MessageCreated", zap.String("channel_id", payload.ChannelID.String()), zap.String("message", payload.PlainText))

	for reg, fun := range mc.mc {
		if regexp.MustCompile(reg).MatchString(payload.PlainText) {
			err := fun(ctx, payload)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (mc *messageCreated) init(bh *BotHandler) {
	mc.mc = make(map[string]func(ctx context.Context, payload payload.EventMessagePayload) error)
	mc.m = sync.Mutex{}
	mc.add(`.*`, bh.echo)
}

func (mc *messageCreated) add(key string, f func(ctx context.Context, payload payload.EventMessagePayload) error) {
	mc.m.Lock()
	defer mc.m.Unlock()

	mc.mc[key] = f
}

func (bh *BotHandler) echo(ctx context.Context, payload payload.EventMessagePayload) error {
	return bh.cl.SendMessage(ctx, payload.ChannelID, payload.PlainText, true)
}
