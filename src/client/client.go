package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/domain"
)

type Client interface {
	SendMessage(ctx context.Context, channelId uuid.UUID, message string, embed bool) error
	AddStamp(ctx context.Context, messageId uuid.UUID, stampID uuid.UUID, count int) error
	GetStampIdByName(ctx context.Context, name string) (uuid.UUID, error)
	GetAllUserIds(ctx context.Context) ([]uuid.UUID, error)
	GetUserInfo(ctx context.Context, userId uuid.UUID) (*domain.TraqUser, error)
}
