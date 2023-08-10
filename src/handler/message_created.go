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

type mentionMessageCreated struct {
	me map[string]func(ctx context.Context, payload payload.EventMessagePayload) error
	m  sync.Mutex
}

var mc messageCreated
var mmc mentionMessageCreated

func (bh *BotHandler) MessageCreatedHandler(ctx context.Context, payload payload.EventMessagePayload) error {
	bh.logger.Debug("MessageCreated", zap.String("channel_id", payload.ChannelID.String()), zap.String("message", payload.PlainText))

	mentioned := false
	for i := range payload.Embedded {
		if e := payload.Embedded[i]; e.Type == "user" && e.ID == botUserId {
			mentioned = true
		}
	}

	if mentioned {
		mmc.m.Lock()
		defer mmc.m.Unlock()
		for reg, f := range mmc.me {
			if regexp.MustCompile(reg).MatchString(payload.PlainText) {
				err := f(ctx, payload)
				if err != nil {
					return err
				}
			}
		}
	} else {
		mc.m.Lock()
		defer mc.m.Unlock()
		for reg, f := range mc.mc {
			if regexp.MustCompile(reg).MatchString(payload.PlainText) {
				err := f(ctx, payload)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (mc *messageCreated) init(bh *BotHandler) {
	mc.mc = make(map[string]func(ctx context.Context, payload payload.EventMessagePayload) error)
	mc.m = sync.Mutex{}
	// mc.add(`.*`, bh.echo)
}

func (mmc *mentionMessageCreated) init(bh *BotHandler) {
	mmc.me = make(map[string]func(ctx context.Context, payload payload.EventMessagePayload) error)
	mmc.m = sync.Mutex{}
	mmc.add(`.*`, bh.echo)
}

func (mc *messageCreated) add(key string, f func(ctx context.Context, payload payload.EventMessagePayload) error) {
	mc.m.Lock()
	defer mc.m.Unlock()

	mc.mc[key] = f
}

func (mmc *mentionMessageCreated) add(key string, f func(ctx context.Context, payload payload.EventMessagePayload) error) {
	mmc.m.Lock()
	defer mmc.m.Unlock()

	mmc.me[key] = f
}

func (bh *BotHandler) echo(ctx context.Context, payload payload.EventMessagePayload) error {
	return bh.cl.SendMessage(ctx, payload.ChannelID, payload.PlainText, true)
}
