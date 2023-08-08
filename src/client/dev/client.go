package dev

import "go.uber.org/zap"

type DevClient struct {
	logger *zap.Logger
}

func NewDevClient(l *zap.Logger) *DevClient {
	l.Info("Dev Client")
	return &DevClient{
		logger: l.Named("Dev Client"),
	}
}
