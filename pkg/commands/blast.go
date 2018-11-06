package commands

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

func stringInList(haystack []string, needle string) bool {
	for _, test := range haystack {
		if test == needle {
			return true
		}
	}
	return false
}

func Blast(c *cli.Context) error {
	var channels []string
	channelFile := c.String("channels-list")
	slackToken := c.GlobalString("slack-token")
	messageFile := c.String("message-file")
	forReal := c.Bool("for-real")

	if channelFile != "" {
		file, err := ioutil.ReadFile(channelFile)
		if err != nil {
			return err
		}
		channels = strings.Split(string(file), "\n")
	}

	api := slack.New(slackToken)

	slackChannels, err := api.GetChannels(true)

	if err != nil {
		return err
	}

	sendChannels := []slack.Channel{}
	names := []string{}
	for _, c := range slackChannels {
		if stringInList(channels, c.Name) {
			sendChannels = append(sendChannels, c)
			names = append(names, c.Name)
		}
	}

	fmt.Printf("the following channels will be notified of your message: \n%s\n", strings.Join(names, "\n"))

	var message []byte
	if messageFile != "" {
		message, err = ioutil.ReadFile(messageFile)
		if err != nil {
			return err
		}
	}

	if forReal {
		for _, c := range sendChannels {
			_, _, err := api.PostMessage(c.ID, string(message), slack.PostMessageParameters{
				Markdown:  true,
				Parse:     "full",
				LinkNames: 1,
			})
			if err != nil {
				fmt.Printf("error sending message to %s - %s", c.Name, err.Error())
			}
		}
	}

	return nil
}
