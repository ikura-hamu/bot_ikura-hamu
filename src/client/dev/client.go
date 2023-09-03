package dev

import (
	"sync"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	iieStampId    uuid.UUID = uuid.MustParse("e03fa967-2c80-1042-b861-cc34d7236dd6")
	yuugenStampId uuid.UUID = uuid.MustParse("636e3dd4-772f-6bd8-3722-69bd5c2eb1a2")
)

type DevClient struct {
	logger *zap.Logger
	stamps map[string]uuid.UUID
	mu     sync.Mutex
}

func NewDevClient(l *zap.Logger) *DevClient {
	l.Info("Dev Client")

	stampsMap := map[string]uuid.UUID{
		"iie":    iieStampId,
		"yuugen": yuugenStampId,
	}

	return &DevClient{
		logger: l.Named("Dev Client"),
		stamps: stampsMap,
	}
}
