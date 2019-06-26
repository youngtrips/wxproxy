package web

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/youngtrips/wxproxy/internal/config"
	"github.com/youngtrips/wxproxy/internal/log"
)

func Run(cfg *config.ServerInfo) {
	e := echo.New()
	e.Debug = true
	if cfg.Http.Mode == "debug" {
		e.Debug = true
	} else if cfg.Http.Mode == "prod" {
		e.Debug = false
	}

	e.Use(middleware.Recover())

	route_init(e)

	addr := fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)
	log.Info("try to listen on : ", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal("start server failed: ", err)
	}
}
