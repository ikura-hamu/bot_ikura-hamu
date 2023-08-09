package client

import (
	"context"

	"github.com/google/uuid"
)

type Client interface {
	SendMessage(ctx context.Context, channelId uuid.UUID, message string, embed bool) error
}
