package router

import (
	"net/http"

	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
	"github.com/ikura-hamu/bot_ikura-hamu/src/client/dev"
	"github.com/ikura-hamu/bot_ikura-hamu/src/client/traq"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/ikura-hamu/bot_ikura-hamu/src/handler"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Setup(logger *zap.Logger, mode conf.Mode) {
	var client client.Client
	switch mode {
	case conf.ProdMode:
		client = traq.NewTraqClient(logger)
	case conf.DevMode:
		client = dev.NewDevClient(logger)
	}

	bh := newBotRouter(handler.NewBotHandler(impl.NewBotRepository(logger), client, logger), logger)

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

	e.POST("/bot", bh.botHandlerFunc, botVerifyMiddleware)

	e.GET("/ping", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	e.Start(":8080")
}
