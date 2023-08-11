package dev

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const STAMP_MSG = "stamp"

var (
	iieStampId    uuid.UUID = uuid.MustParse("e03fa967-2c80-1042-b861-cc34d7236dd6")
	yuugenStampId uuid.UUID = uuid.MustParse("636e3dd4-772f-6bd8-3722-69bd5c2eb1a2")
)

func (dc *DevClient) AddStamp(ctx context.Context, messageId uuid.UUID, stampId uuid.UUID, count int) error {
	dc.logger.Debug(STAMP_MSG,
		zap.String("message", messageId.String()),
		zap.String("stampId", stampId.String()),
		zap.Int("stamp count", count))
	return nil
}

func (dc *DevClient) GetAllStamps(ctx context.Context) (map[string]uuid.UUID, error) {
	stampsMap := map[string]uuid.UUID{
		"iie":    iieStampId,
		"yuugen": yuugenStampId,
	}
	return stampsMap, nil
}
