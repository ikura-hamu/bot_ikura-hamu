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
	channelId := payload.MessagePayload.ChannelID
	if err := bh.cl.JoinChannel(ctx, channelId); err != nil {
		return err
	}
	if _, err := bh.cl.SendMessage(ctx, channelId, ":oisu-:", false); err != nil {
		return err
	}
	return nil
}

func (bh *BotHandler) leave(ctx context.Context, payload payload.EventMessagePayload) error {
	channelId := payload.MessagePayload.ChannelID
	if err := bh.cl.LeaveChannel(ctx, channelId); err != nil {
		return err
	}
	if _, err := bh.cl.SendMessage(ctx, channelId, ":wave:", false); err != nil {
		return err
	}
	return nil
}
