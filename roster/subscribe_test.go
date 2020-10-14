package roster

import (
	"context"
	"fmt"
	"github.com/slack-go/slack"
	"testing"
)

func TestFilter(t *testing.T) {
	a, b,  c := server().Subscribe(context.Background(), slack.SlashCommand{
		Token:          "asdasd",
		TeamID:         "asdasd",
		TeamDomain:     "asdasd",
		EnterpriseID:   "asdasd",
		EnterpriseName: "asdasd",
		ChannelID:      "asdasd",
		ChannelName:    "foobar",
		UserID:         "asdasd",
		UserName:       "asdasd",
		Command:        "asdasd",
		Text:           `add "9 * * * *" "THis is a message" @joshuacarpeggiani `,
		ResponseURL:    "",
		TriggerID:      "",
	})
	fmt.Println(a, b, c)
}

func server() Server {
	return NewServer("slack", "https://us-central1-joshcarp-installer.cloudfunctions.net/respond", "joshcarp-installer", "asdasd", "asdasd")
}

