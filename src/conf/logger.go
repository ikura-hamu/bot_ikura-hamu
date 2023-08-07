package conf

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	mode := "prod"
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}

	switch mode {
	case "dev":
		level := zap.NewAtomicLevel()
		level.SetLevel(zapcore.DebugLevel)

		myConfig := zap.Config{
			Level:    level,
			Encoding: "console",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "Time",
				LevelKey:       "Level",
				NameKey:        "Name",
				CallerKey:      "Caller",
				MessageKey:     "Msg",
				StacktraceKey:  "St",
				EncodeLevel:    zapcore.CapitalColorLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		logger, err = myConfig.Build()
		if err != nil {
			return nil, err
		}
	default:
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	}

	logger.Info("zap", zap.String("mode", mode))

	return logger, nil
}
