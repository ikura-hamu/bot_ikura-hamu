package router

import (
	"github.com/ikura-hamu/bot_ikura-hamu/src/handler"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Setup(logger *zap.Logger) {
	bh := newBotRouter(*handler.NewBotHandler(impl.NewBotRepository(logger), logger), logger)

	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	e.GET("/bot", bh.botHandlerFunc)

	e.Start(":8080")
}
