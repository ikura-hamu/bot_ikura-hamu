package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/model"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (br *BotRepository) CreateBioQuiz(ctx context.Context, channelId uuid.UUID, messageId uuid.UUID, answer string) error {
	_, err := br.db.Collection("bio_quiz").InsertOne(ctx, model.BioQuiz{
		ChannelId:  channelId,
		MessageId:  messageId,
		Answer:     answer,
		IsAnswered: false,
	})
	if err != nil {
		return handleError(err)
	}
	return nil
}

type searchBioQuiz struct {
	ChannelId  uuid.UUID `bson:"channel_id"`
	IsAnswered bool      `bson:"is_answered"`
}

func (br *BotRepository) GetNotAnsweredBioQuiz(ctx context.Context, channelId uuid.UUID) (*model.BioQuiz, error) {
	var bioQuiz model.BioQuiz
	err := br.db.
		Collection("bio_quiz").
		FindOne(ctx, searchBioQuiz{ChannelId: channelId, IsAnswered: false}).
		Decode(&bioQuiz)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository.ErrBioQuizNotFound
	}
	if err != nil {
		return nil, handleError(err)
	}

	return &bioQuiz, nil
}

func (br *BotRepository) AnswerBioQuiz(ctx context.Context, id string) error {
	br.logger.Debug("AnswerBioQuiz", zap.String("id", id))
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return handleError(err)
	}
	result, err := br.db.
		Collection("bio_quiz").
		UpdateByID(ctx, objectId, map[string]interface{}{"$set": map[string]interface{}{"is_answered": true}})
	if err != nil {
		return handleError(err)
	}

	br.logger.Debug("AnswerBioQuiz", zap.Any("result", result))
	return nil
}
