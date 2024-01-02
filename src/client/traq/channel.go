package traq

import (
	"context"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"

	"github.com/google/uuid"
	"github.com/traPtitech/go-traq"
)

var botId = conf.GetBotId()

func (tc *TraqClient) JoinChannel(ctx context.Context, channelId uuid.UUID) error {
	req := traq.PostBotActionJoinRequest{ChannelId: channelId.String()}
	_, err := tc.client.BotApi.LetBotJoinChannel(ctx, botId).PostBotActionJoinRequest(req).Execute()
	if err != nil {
		return handleError(err)
	}
	return nil
}

func (tc *TraqClient) LeaveChannel(ctx context.Context, channelId uuid.UUID) error {
	req := traq.PostBotActionLeaveRequest{ChannelId: channelId.String()}
	_, err := tc.client.BotApi.LetBotLeaveChannel(ctx, botId).PostBotActionLeaveRequest(req).Execute()
	if err != nil {
		return handleError(err)
	}
	return nil
}
