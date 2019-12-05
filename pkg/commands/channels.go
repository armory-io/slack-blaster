package commands

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/ryanuber/go-glob"
)

func Channels(slackToken string, pattern string, sharedOnly bool) error {
	api := slack.New(slackToken)

	slackChannels := []slack.Channel{}
	types := []string{"public_channel"}
	err := error(nil)
	cursor := ""

	// Loop through pagination till we get all slack channels.
	for hasMorePages := true; hasMorePages; hasMorePages = !(cursor == "") {
		slackChannelsPage := []slack.Channel{}
		params := slack.GetConversationsParameters{cursor, "true", 1000, types}

		slackChannelsPage, cursor, err = api.GetConversations(&params)

		if err != nil {
			return err
		}

		slackChannels = append(slackChannels, slackChannelsPage...)
	}

	names := []string{}

	for _, c := range slackChannels {
		// If a pattern was specified ensure the channel name matches.
		if !glob.Glob(pattern, c.Name) {
			continue
		}
		// If looking for shared only channels check the property.
		if sharedOnly && !c.IsShared {
			continue
		}

		names = append(names, c.Name)
	}

	fmt.Printf("channels matching %s\n%s\n", pattern, strings.Join(names, "\n"))
	return nil
}
