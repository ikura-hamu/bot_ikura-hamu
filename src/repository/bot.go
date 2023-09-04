package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/model"
)

type BotRepository interface {
	CreateBioQuiz(ctx context.Context, channelId uuid.UUID, messageId uuid.UUID, answer string) error
	GetNotAnsweredBioQuiz(ctx context.Context, channelId uuid.UUID) (*model.BioQuiz, error)
	AnswerBioQuiz(ctx context.Context, channelId uuid.UUID) error
}
