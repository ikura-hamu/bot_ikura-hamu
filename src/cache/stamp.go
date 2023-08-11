package cache

import "github.com/google/uuid"

type StampCache interface {
	GetStampIdByName(name string) (uuid.UUID, bool, error)
	SetStampCache(data map[string]uuid.UUID) error
}