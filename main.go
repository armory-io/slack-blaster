package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/nlopes/slack"
)

func stringInList(haystack []string, needle string) bool {
	for _, test := range haystack {
		if test == needle {
			return true
		}
	}
	return false
}

var wg sync.WaitGroup

func main() {
	// pattern := flag.String("pattern", "*", "channel name filter")
	messageFile := flag.String("message-file", "", "file containing message content")
	channelFile := flag.String("channels-file", "", "file containing a list of channels to blast")
	slackToken := flag.String("slack-token", "", "slack api token")
	forReal := flag.Bool("for-real", false, "send the message...for real")
	// flag for a file containing channel names
	flag.Parse()

	if *slackToken == "" {
		token := os.Getenv("SLACK_TOKEN")
		slackToken = &token
	}

	var channels []string
	if *channelFile != "" {
		file, err := ioutil.ReadFile(*channelFile)
		if err != nil {
			panic(err)
		}
		channels = strings.Split(string(file), "\n")
	}

	api := slack.New(*slackToken)

	slackChannels, err := api.GetChannels(true)

	if err != nil {
		panic(err)
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
	if *messageFile != "" {
		message, err = ioutil.ReadFile(*messageFile)
		if err != nil {
			panic(err)
		}
	}

	if *forReal {
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
				fmt.Printf("Done sending message to %s", ch.Name)
				wg.Done()
			}(&cx, &wg)
		}
		wg.Wait()
	}

}
