package traq

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/go-traq"
)

func (tc *TraqClient) SendMessage(ctx context.Context, channelId uuid.UUID, message string, embed bool) (uuid.UUID, error) {
	var req traq.PostMessageRequest
	req.SetContent(message)
	req.SetEmbed(embed)
	mes, _, err := tc.client.MessageApi.PostMessage(ctx, channelId.String()).PostMessageRequest(req).Execute()
	if err != nil {
		return uuid.Nil, handleError(err)
	}
	return uuid.MustParse(mes.Id), nil
}
