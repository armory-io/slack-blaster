package main

import (
	"fmt"
	"os"

	"github.com/armory-io/slack-blaster/pkg/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "slack-token", EnvVar: "SLACK_TOKEN", Usage: "Slack token to use when connecting to the API"},
	}
	app.Commands = []cli.Command{
		{
			Name:        "channels",
			Description: "list and filter channel names",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "pattern", Usage: "pattern to use when matching channel name", Value: "*"},
			},
			Action: func(c *cli.Context) error {
				if err := commands.Channels(c); err != nil {
					fmt.Println(err.Error())
					return err
				}
				return nil
			},
		},
		{
			Name:        "blast",
			Description: "bulk send message",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "channels-list", Usage: "list of channels to blast"},
				cli.StringFlag{Name: "message-file", Usage: "file containing the message to send"},
				cli.BoolFlag{Name: "for-real", Usage: "send the message...for real"},
			},
			Action: func(c *cli.Context) error {
				if err := commands.Blast(c); err != nil {
					fmt.Println(err.Error())
					return err
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
