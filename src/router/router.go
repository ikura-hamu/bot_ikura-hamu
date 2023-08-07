package router

import (
	"github.com/ikura-hamu/bot_ikura-hamu/src/handler"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository/impl"
	"github.com/labstack/echo/v4"
)

func Setup() {
	bh := newBotRouter(*handler.NewBotHandler(impl.NewBotRepository()))

	e := echo.New()
	e.GET("/bot", bh.botHandlerFunc)

	e.Start(":8080")
}
