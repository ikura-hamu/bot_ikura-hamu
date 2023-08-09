package router

import (
	"net/http"

	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/labstack/echo/v4"
)

func botVerifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-TOKEN")]) == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "no bot token")
		} else if c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-TOKEN")][0] != conf.GetBotToken() {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		return next(c)
	}
}
