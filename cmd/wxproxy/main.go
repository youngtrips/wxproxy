package main

import (
	"github.com/youngtrips/wxproxy/internal/appkit"
	"github.com/youngtrips/wxproxy/server"
)

const (
	APP_NAME  = "wxproxy"
	APP_USAGE = "start a wxproxy server instance"
)

var APP_VERSION = "unknown"

func main() {
	appkit.New(APP_NAME, APP_USAGE, APP_VERSION, server.Run)
}
