package roster

import (
	"testing"
)

func TestFilter2(t *testing.T) {

	//server().Unsubscribe(cmd)
}

func TestFilter(t *testing.T) {
	//cmd := slack.SlashCommand{
	//	Token:          "asdasd",
	//	TeamID:         "aaaaa",
	//	TeamDomain:     "asdasd",
	//	EnterpriseID:   "asdasd",
	//	EnterpriseName: "asdasd",
	//	ChannelID:      "Foobar",
	//	ChannelName:    "Blah_Blah",
	//	UserID:         "asdasd",
	//	UserName:       "asdasd",
	//	Command:        "asdasd",
	//	Text:           `add "* * * * *" "THis is a message" @joshuacarpeggiani `,
	//	ResponseURL:    "",
	//	TriggerID:      "",
	//}
	//a, b,  c := server().Subscribe(context.Background(), cmd)
	//fmt.Println(a, b, c)
	//server().Respond(context.Background(), time.Now())
}

func server() Server {
	return NewServer("", "joshcarp-installer", "", "")
}

