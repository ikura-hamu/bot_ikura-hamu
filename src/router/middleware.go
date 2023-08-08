package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func botVerifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-TOKEN")]) == 0 ||
			c.Request().Header[http.CanonicalHeaderKey("X-TRAQ-BOT-TOKEN")][0] != "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "no bot token")
		}
		return next(c)
	}
}
