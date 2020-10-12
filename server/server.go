package main

import (
	"fmt"
	"github.com/joshcarp/rosterbot"
	"github.com/slack-go/slack"
)

func main() {
	a, b := rosterbot.SlackCommandSubscribe(slack.SlashCommand{
		Token:          "",
		TeamID:         "",
		TeamDomain:     "",
		EnterpriseID:   "",
		EnterpriseName: "",
		ChannelID:      "",
		ChannelName:    "",
		UserID:         "",
		UserName:       "",
		Command:        "",
		Text:           `"* * * * *" "foobar" @joshuacarpeggiani @joshuacarpeggiani`,
		ResponseURL:    "",
		TriggerID:      "",
	})
	fmt.Println(a,b)
}
