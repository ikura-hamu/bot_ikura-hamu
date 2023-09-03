package router

import (
	"fmt"
	"net/http"

	"github.com/ikura-hamu/bot_ikura-hamu/pkg/payload"
	"github.com/ikura-hamu/bot_ikura-hamu/src/handler"
	"github.com/labstack/echo/v4"
	traqbot "github.com/traPtitech/traq-bot"
	"go.uber.org/zap"
)

type botRouter struct {
	bh     handler.BotHandler
	logger *zap.Logger
}

func newBotRouter(bh handler.BotHandler, l *zap.Logger) *botRouter {
	return &botRouter{
		bh:     bh,
		logger: l,
	}
}

func (br *botRouter) botHandlerFunc(c echo.Context) error {
	var err error
	eventHeader := c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-EVENT")]
	if len(eventHeader) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "no X-TRAQ-BOT-EVENT header")
	}
	event := c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-EVENT")][0]
	switch event {
	case traqbot.Ping:
		return c.NoContent(http.StatusNoContent)
	case traqbot.MessageCreated:
		var body payload.EventMessagePayload
		err = c.Bind(&body)
		if err != nil {
			br.logger.Debug("bad request body", zap.Error(err))
			return echo.NewHTTPError(http.StatusBadRequest, "bad request body")
		}
		err = br.bh.MessageCreatedHandler(c.Request().Context(), body)
	default:
		return echo.NewHTTPError(http.StatusNotImplemented, fmt.Sprintf("event '%s' is not implemented", event))
	}

	if err != nil {
		br.logger.Error("error", zap.Error(err))
	}

	return c.NoContent(http.StatusOK)
}
