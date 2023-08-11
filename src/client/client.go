package client

import (
	"context"

	"github.com/google/uuid"
)

type Client interface {
	SendMessage(ctx context.Context, channelId uuid.UUID, message string, embed bool) error
	AddStamp(ctx context.Context, messageId uuid.UUID, stampID uuid.UUID, count int) error
	GetAllStamps(ctx context.Context) (map[string]uuid.UUID, error)
}
