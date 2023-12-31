package handler

import (
	"context"
	"github.com/ikura-hamu/bot_ikura-hamu/pkg/payload"
	"regexp"
)

var (
	joinMessageReg  = regexp.MustCompile(`join\s*$`)
	leaveMessageReg = regexp.MustCompile(`leave\s*$`)
)

func (bh *BotHandler) join(ctx context.Context, payload payload.EventMessagePayload) error {
	message := payload.MessagePayload.PlainText
	channelId := payload.MessagePayload.ChannelID
	if joinMessageReg.MatchString(message) {
		return bh.cl.JoinChannel(ctx, channelId)
	}
	return nil
}

func (bh *BotHandler) leave(ctx context.Context, payload payload.EventMessagePayload) error {
	message := payload.MessagePayload.PlainText
	channelId := payload.MessagePayload.ChannelID
	if leaveMessageReg.MatchString(message) {
		return bh.cl.LeaveChannel(ctx, channelId)
	}
	return nil
}
