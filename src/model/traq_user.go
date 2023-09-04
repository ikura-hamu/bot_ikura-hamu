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

func (tu *TraqUser) GetId() uuid.UUID {
	return tu.id
}

func (tu *TraqUser) GetName() string {
	return tu.name
}

func (tu *TraqUser) GetDisplayName() string {
	return tu.displayName
}

func (tu *TraqUser) GetBio() string {
	return tu.bio
}

func (tu *TraqUser) GetLastOnline() time.Time {
	return tu.lastOnline
}
