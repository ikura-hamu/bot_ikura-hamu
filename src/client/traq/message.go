package traq

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/go-traq"
)

func (tc *TraqClient) SendMessage(ctx context.Context, channelId uuid.UUID, message string, embed bool) error {
	var req traq.PostMessageRequest
	req.SetContent(message)
	req.SetEmbed(embed)
	_, _, err := tc.client.MessageApi.PostMessage(ctx, channelId.String()).PostMessageRequest(req).Execute()
	if err != nil {
		return handleError(err)
	}
	return nil
}
