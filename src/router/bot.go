package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ikura-hamu/bot_ikura-hamu/src/handler"
	"github.com/labstack/echo/v4"
	traqbot "github.com/traPtitech/traq-bot"
)

type botRouter struct {
	bh handler.BotHandler
}

func newBotRouter(bh handler.BotHandler) *botRouter {
	return &botRouter{
		bh: bh,
	}
}

func (br *botRouter) botHandlerFunc(c echo.Context) error {
	var err error
	event := c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-EVENT")][0]
	switch event {
	case traqbot.Ping:
		return c.NoContent(http.StatusNoContent)
	case traqbot.MessageCreated:
		err = br.bh.MessageCreatedHandler(c.Request().Context(), traqbot.MessageCreatedPayload{})
	default:
		return echo.NewHTTPError(http.StatusNotImplemented, fmt.Sprintf("event '%s' is not implemented", event))
	}

	if err != nil {
		log.Printf("error: %v", err)
	}

	return c.NoContent(http.StatusOK)
}
