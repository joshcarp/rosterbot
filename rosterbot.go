package rosterbot

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/slack-go/slack"
)

type server struct {
	*slack.Client
}

type RosterPayload struct {
	command
	Channel string
}

func (r RosterPayload) toMap() map[string]string {
	return map[string]string{
		"channel":   r.Channel,
		"starttime": r.command.StartTime.String(),
		"message":   r.command.Message,
		"users":     strings.Join(r.command.Users, ", "),
	}
}

func (r *RosterPayload) fromMap(m map[string]string) {
	r.Channel = m["channel"]
	r.StartTime, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", m["starttime"])
	r.Message = m["message"]
	r.Users = strings.Split(m["users"], ", ")
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := server{Client: nil}
	s.ServeHTTP(w, r)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	cmd, _ := slack.SlashCommandParse(r)
	_, err := SlackCommandSubscribe(cmd)
	if err != nil {
		log.Println("{" + err.Error() + "}")
	} else {
		log.Println("yay")
	}

}

func SlackCommandSubscribe(cmd slack.SlashCommand) (*pubsub.Subscription, error) {
	rosterbotCommand, err := ParseCommand(cmd.Text)
	if err != nil{
		return nil, err
	}
	payload := RosterPayload{command: rosterbotCommand, Channel: cmd.ChannelID}
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, "joshcarp-installer")
	if err != nil{
		return nil, err
	}
	return pubsubService.CreateSubscription(ctx, payload.Channel, pubsub.SubscriptionConfig{
		Topic:  pubsubService.Topic("slack"),
		PushConfig: pubsub.PushConfig{
			Endpoint:   os.Getenv("PUSH_URL"),
			Attributes: payload.toMap(),
		},
	})
}

func (s *server) SlackRespond(w http.ResponseWriter, r *http.Request) {
	//log.Println(r)
	//rosterbotCommand, err := ParseCommand(command.Command)
	//pubsubService, _ := pubsub.NewClient(context.Background(), "joshcarp-installer")
	//pubsubService.Subscription("foobar").
	//s.PostMessage(command.ChannelID, slack.MsgOptionText(fmt.Sprintf(`{
	//"channel": "%s",
	//"text": "The time in %s is %s",
	//"as_user": true
	//}`, command.ChannelID, command.Text, "t"), false))
}
