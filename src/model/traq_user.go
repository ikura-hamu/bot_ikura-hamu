package model

import (
	"time"

	"github.com/google/uuid"
)

type TraqUser struct {
	id          uuid.UUID
	name        string
	displayName string
	bio         string
	lastOnline  time.Time
}

func NewTraqUser(id uuid.UUID, name string, displayName string, bio string, lastOnline time.Time) *TraqUser {
	return &TraqUser{
		id:          id,
		name:        name,
		displayName: displayName,
		bio:         bio,
		lastOnline:  lastOnline,
	}
}
