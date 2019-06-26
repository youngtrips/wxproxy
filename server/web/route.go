package web

import (
	"github.com/labstack/echo/v4"
	"github.com/youngtrips/wxproxy/server/web/controller/v1/wx"
)

func route_init(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/wx", wx.ValidateHandler)
	v1.POST("/wx", wx.MessageHandler)
}
