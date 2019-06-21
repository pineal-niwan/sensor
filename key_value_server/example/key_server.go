package main

import (
	"github.com/google/martian/log"
	"github.com/pineal-niwan/sensor/key_value_server/server"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.App{
		Name:`keyServer`,
		Usage: `key - value server`,
		Version:`0.1`,
		Flags:[]cli.Flag{
			&cli.StringFlag{
				Name:`address`,
			},
			&cli.IntFlag{
				Name:`size`,
			},
		},
		Action:serverRun,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Errorf(`run error:%+v`, err)
	}
}

func serverRun(c *cli.Context) error {
	return server.StartKServer(c.String(`address`), c.Int(`size`))
}
