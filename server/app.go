package server

import (
	"context"

	"github.com/youngtrips/wxproxy/internal/config"
	"github.com/youngtrips/wxproxy/server/web"
)

func Run(ctx context.Context) {
	cfg := config.Config()
	web.Run(&cfg)
}
