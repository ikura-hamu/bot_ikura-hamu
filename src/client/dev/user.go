package dev

import (
	"context"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/domain"
)

var (
	userId1 = uuid.MustParse("7265b13d-9e06-42f6-98e3-41ea742f8fb2")
)

func (dc *DevClient) GetAllUserIds(ctx context.Context) ([]uuid.UUID, error) {
	return []uuid.UUID{userId1}, nil
}

func (dc *DevClient) GetUserInfo(ctx context.Context, userId uuid.UUID) (*domain.TraqUser, error) {
	return domain.NewTraqUser(userId, "ikura-hamu", "いくら・はむ", "ひとこと"), nil
}
