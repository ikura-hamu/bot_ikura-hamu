package model

import "github.com/google/uuid"

type BioQuiz struct {
	ChannelId  uuid.UUID `bson:"channel_id"`
	MessageId  uuid.UUID `bson:"message_id"`
	Answer     string    `bson:"answer"`
	IsAnswered bool      `bson:"is_answered"`
}
