package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Run() error {
	app := cli.App{
		Commands: []*cli.Command{&cli.Command{
			Name:    "startServer",
			Aliases: []string{"s", "serve"},
			Usage:   "start http server",
			Action:  startServer,
		}},
	}

	return app.Run(os.Args)
}
