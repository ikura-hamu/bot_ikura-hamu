package model

import "github.com/google/uuid"

type BioQuiz struct {
	Id         string    `bson:"_id,omitempty"`
	ChannelId  uuid.UUID `bson:"channel_id"`
	MessageId  uuid.UUID `bson:"message_id,omitempty"`
	Answer     string    `bson:"answer,omitempty"`
	IsAnswered bool      `bson:"is_answered"`
}
