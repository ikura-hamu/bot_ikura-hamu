package bigcache

import (
	"context"
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"go.uber.org/zap"
)

type StampCache struct {
	cache  *bigcache.BigCache
	logger *zap.Logger
}

func NewStampCache(l *zap.Logger) (*StampCache, error) {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*time.Hour))
	if err != nil {
		return nil, err
	}
	return &StampCache{cache: cache, logger: l.Named("bigcache")}, nil
}

func (sc *StampCache) GetStampIdByName(name string) (uuid.UUID, bool, error) {
	b, err := sc.cache.Get(name)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return uuid.Nil, false, nil
	}
	if err != nil {
		return uuid.Nil, false, err
	}

	id, err := uuid.FromBytes(b)
	if err != nil {
		return uuid.Nil, false, err
	}
	return id, true, nil
}

func (sc *StampCache) SetStampCache(stampMap map[string]uuid.UUID) error {
	var mError error
	for stampName := range stampMap {
		b, err := stampMap[stampName].MarshalBinary()
		if err != nil {
			mError = multierror.Append(err, mError)
			sc.logger.Error("uuid to byte[]", zap.Error(err))
		}
		err = sc.cache.Set(stampName, b)
		if err != nil {
			mError = multierror.Append(err, mError)
			sc.logger.Error("set cache", zap.Error(err))
		}
	}
	return mError
}
