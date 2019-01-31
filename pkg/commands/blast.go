package commands

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

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
	var wg sync.WaitGroup
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
		wg.Add(len(sendChannels))
		for c := range sendChannels {
			cx := sendChannels[c]
			go func(chanl *slack.Channel, wg *sync.WaitGroup) {
				// join the channel first if we aren't in it yet
				ch, err := api.JoinChannel(chanl.Name)
				if err != nil {
					fmt.Println(err)
					wg.Done()
					return
				}
				fmt.Printf("joined channel %s\n", ch.Name)
				_, _, err = api.PostMessage(ch.ID, string(message), slack.PostMessageParameters{
					Markdown:  true,
					Parse:     "full",
					LinkNames: 1,
				})
				if err != nil {
					fmt.Printf("error sending message to %s - %s", ch.Name, err.Error())
				}
				fmt.Printf("Done sending message to %s\n", ch.Name)
				wg.Done()
			}(&cx, &wg)
		}
		wg.Wait()
	}

	return nil
}
