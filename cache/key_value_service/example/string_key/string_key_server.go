package main

import (
	"github.com/google/martian/log"
	"github.com/pineal-niwan/sensor/cache/key_value_service/server"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.App{
		Name:    `string key server`,
		Usage:   `key - value server`,
		Version: `0.1`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: `address`,
			},
			&cli.IntFlag{
				Name: `size`,
			},
		},
		Action: stringKeyServerRun,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Errorf(`run error:%+v`, err)
	}
}

func stringKeyServerRun(c *cli.Context) error {
	return server.StartStringKeyServer(c.String(`address`), c.Int(`size`))
}
