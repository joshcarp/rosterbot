package roster

import (
	"testing"
)

func TestFilter(t *testing.T) {
	//cmd := slack.SlashCommand{
	//	Token:          "asdasd",
	//	TeamID:         "aaaaa",
	//	TeamDomain:     "asdasd",
	//	EnterpriseID:   "asdasd",
	//	EnterpriseName: "asdasd",
	//	ChannelID:      "aaaa",
	//	ChannelName:    "foobar",
	//	UserID:         "asdasd",
	//	UserName:       "asdasd",
	//	Command:        "asdasd",
	//	Text:           `add "9 * * * *" "THis is a message" @joshuacarpeggiani `,
	//	ResponseURL:    "",
	//	TriggerID:      "",
	//}
	//a, b,  c := server().Subscribe(context.Background(), cmd)
	//fmt.Println(a, b, c)
	//server().Unsubscribe(cmd)
}

func server() Server {
	return NewServer("https://us-central1-joshcarp-installer.cloudfunctions.net/respond", "joshcarp-installer", "asdasd", "asdasd")
}

