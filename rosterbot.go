package rosterbot

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/slack-go/slack"
	"google.golang.org/api/pubsub/v1"
)

type server struct {
	*slack.Client
}

type RosterPayload struct{
	command
	Channel string
}

func (r RosterPayload)toMap()map[string]string{
	return map[string]string{
		"channel": r.Channel,
		"starttime": r.command.StartTime.String(),
		"message": r.command.Message,
		"users": strings.Join(r.command.Users, ", "),
	}
}

func (r *RosterPayload)fromMap(m map[string]string){
	r.Channel = m["channel"]
	r.StartTime, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", m["starttime"])
	r.Message = m["message"]
	r.Users = strings.Split(m["users"], ", ")
}

func ServeHTTP(w http.ResponseWriter, r *http.Request){
	s := server{Client: nil}
	s.ServeHTTP(w, r)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	cmd, _ := slack.SlashCommandParse(r)
	_, err := SlackCommandSubscribe(cmd)
	if err != nil{
		log.Println(err)
	}

}

func SlackCommandSubscribe(cmd slack.SlashCommand) (*pubsub.Subscription, error) {
	rosterbotCommand, _ := ParseCommand(cmd.Command)
	payload := RosterPayload{command: rosterbotCommand, Channel: cmd.Command}
	ctx := context.Background()
	pubsubService, _ := pubsub.NewService(ctx)
	return pubsubService.Projects.Subscriptions.Create(
		payload.Channel,
		&pubsub.Subscription{
			Topic:"projects/joshcarp-installer/topics/slack",
			PushConfig: &pubsub.PushConfig{
				Attributes:      payload.toMap(),
				PushEndpoint:    os.Getenv("PUSH_URL"),
			}}).Do()
}

func (s *server) SlackRespond(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	//var res *http.Response
	//var err error
	//command, err := slack.SlashCommandParse(r)
	//rosterbotCommand, err := ParseCommand(command.Command)

	//s.PostMessage(command.ChannelID, slack.MsgOptionText(fmt.Sprintf(`{
	//"channel": "%s",
	//"text": "The time in %s is %s",
	//"as_user": true
	//}`, command.ChannelID, command.Text, "t"), false))
}
