package dev

import (
	"context"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
	"go.uber.org/zap"
)

const STAMP_MSG = "stamp"

func (dc *DevClient) AddStamp(ctx context.Context, messageId uuid.UUID, stampId uuid.UUID, count int) error {
	dc.logger.Debug(STAMP_MSG,
		zap.String("message", messageId.String()),
		zap.String("stampId", stampId.String()),
		zap.Int("stamp count", count))
	return nil
}

func (dc *DevClient) GetAllStamps(ctx context.Context) (map[string]uuid.UUID, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	stampsMap := dc.stamps
	return stampsMap, nil
}

func (dc *DevClient) GetStampIdByName(ctx context.Context, name string) (uuid.UUID, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	id, ok := dc.stamps[name]
	if !ok {
		return uuid.Nil, client.ErrInvalidStampName
	}
	return id, nil
}
