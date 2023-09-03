package domain

import (
	"github.com/google/uuid"
)

type TraqUser struct {
	id          uuid.UUID
	name        string
	displayName string
	bio         string
}

func NewTraqUser(id uuid.UUID, name string, displayName string, bio string) *TraqUser {
	return &TraqUser{
		id:          id,
		name:        name,
		displayName: displayName,
		bio:         bio,
	}
}
