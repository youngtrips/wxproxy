package appkit

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/youngtrips/wxproxy/internal/config"
	"github.com/youngtrips/wxproxy/internal/log"
	"gopkg.in/urfave/cli.v2"
)

func New(name string, usage string, version string, loop func(ctx context.Context)) {
	binpath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	app := &cli.App{
		Name:    name,
		Usage:   usage,
		Version: version,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: filepath.Join(binpath, "..", fmt.Sprintf("conf/%s.yml", name)),
				Usage: "specified config file",
			},
			&cli.BoolFlag{
				Name:  "pprof",
				Usage: "start profiling server on :6060",
			},
		},
		Action: func(c *cli.Context) error {
			if err := config.Load(c.String("config")); err != nil {
				panic("load config failed: " + c.String("config"))
			}

			// log init
			log.Init(name)
			defer log.Sync()

			// c.Bool("pprof")
			if c.Bool("pprof") {
				go http.ListenAndServe(":6060", nil)
			}

			ctx, cancelFunc := context.WithCancel(context.TODO())
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, os.Kill, syscall.SIGTERM, os.Interrupt, syscall.SIGINT, syscall.SIGHUP)
			log.Infof("start server [%s]", name)
			go loop(ctx)
			select {
			case <-sc:
				cancelFunc()
				log.Infof("stop server [%s]", name)
			}
			return nil
		},
	}
	app.Run(os.Args)
}
