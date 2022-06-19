package main

import (
	"log"
	"os"

	"github.com/telegram-bot-template/bot"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config-file",
				Aliases: []string{"cf"},
				Value:   "",
				Usage:   "The config file path of your program",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run this app",
				Action: func(c *cli.Context) error {
					path := ""
					if c.String("config-file") != "" {
						path = c.String("config-file")
					} else {
						path = "config.yaml"
					}
					bot.ConnectMyBot(path)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
