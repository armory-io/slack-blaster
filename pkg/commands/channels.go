package commands

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/ryanuber/go-glob"
	"github.com/urfave/cli"
)

func Channels(c *cli.Context) error {
	pattern := c.String("pattern")
	slackToken := c.GlobalString("slack-token")
	api := slack.New(slackToken)

	slackChannels, err := api.GetChannels(true)

	if err != nil {
		return err
	}

	names := []string{}

	for _, c := range slackChannels {
		if glob.Glob(pattern, c.Name) {
			names = append(names, c.Name)
		}
	}

	fmt.Printf("channels matching %s\n%s", pattern, strings.Join(names, "\n"))
	return nil
}
