package impl

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/model"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func (br *BotRepository) CreateBioQuiz(ctx context.Context, channelId uuid.UUID, messageId uuid.UUID, answer string) error {
	_, err := br.db.Collection("bio_quiz").InsertOne(ctx, model.BioQuiz{
		ChannelId:  channelId,
		MessageId:  messageId,
		Answer:     answer,
		IsAnswered: false,
	})
	if err != nil {
		return err
	}
	return nil
}

func (br *BotRepository) GetNotAnsweredBioQuiz(ctx context.Context, channelId uuid.UUID) (*model.BioQuiz, error) {
	var bioQuiz model.BioQuiz
	err := br.db.
		Collection("bio_quiz").
		FindOne(ctx, model.BioQuiz{ChannelId: channelId, IsAnswered: false}).
		Decode(&bioQuiz)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository.ErrBioQuizNotFound
	}
	if err != nil {
		return nil, err
	}
	return &bioQuiz, nil
}

func (br *BotRepository) AnswerBioQuiz(ctx context.Context, channelId uuid.UUID) error {
	_, err := br.db.
		Collection("bio_quiz").
		UpdateOne(ctx, model.BioQuiz{ChannelId: channelId}, model.BioQuiz{IsAnswered: true})
	if err != nil {
		return err
	}
	return nil
}
