package handler

import (
	"context"
	"regexp"
	"sync"

	"github.com/ikura-hamu/bot_ikura-hamu/pkg/payload"
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

	mc.add(`(にゃん|ニャン)`, bh.cat)
	mc.add(`無限`, bh.infinity)
}

func (mmc *mentionMessageCreated) init(bh *BotHandler) {
	mmc.me = make(map[string]func(ctx context.Context, payload payload.EventMessagePayload) error)
	mmc.m = sync.Mutex{}
	mmc.add(`ひとことクイズ`, bh.bioQuiz)
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

func (bh *BotHandler) cat(ctx context.Context, payload payload.EventMessagePayload) error {
	_, err := bh.cl.SendMessage(ctx, payload.ChannelID,
		"ふぁぼってにゃん♡体操いくよー！ かっわいい私をふぁぼってにゃん♪にゃん！ 純情過ぎててふぁぼってにゃん♪にゃん！ テンション高くてふぁぼってにゃん♪にゃん！ 性格良すぎてふぁぼってにゃん♪ ふぁぼってにゃ〜ぁ〜ん♪ふぁぼってにゃ〜ぁ〜ん♪ マジでふぁぼってにゃんにゃんにゃ〜ん♪",
		true)
	return err
}

func (bh *BotHandler) infinity(ctx context.Context, payload payload.EventMessagePayload) error {
	iieStampId, err := bh.cl.GetStampIdByName(ctx, "iie")
	if err != nil {
		return err
	}
	err = bh.cl.AddStamp(ctx, payload.ID, iieStampId, 1)
	if err != nil {
		return err
	}

	finiteStampId, err := bh.cl.GetStampIdByName(ctx, "yuugen")
	if err != nil {
		return err
	}
	err = bh.cl.AddStamp(ctx, payload.ID, finiteStampId, 1)
	if err != nil {
		return err
	}
	return nil
}
