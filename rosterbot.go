package rosterbot

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joshcarp/rosterbot/cron"
	"github.com/joshcarp/rosterbot/filter"

	"cloud.google.com/go/pubsub"
	"github.com/slack-go/slack"
)

type server struct {
	*slack.Client
}

type RosterPayload struct {
	command
	Channel string
	Token string
}

func (r RosterPayload) toMap() map[string]string {
	return map[string]string{
		"channel":   r.Channel,
		"starttime": r.command.StartTime.String(),
		"message":   r.command.Message,
		"users":     strings.Join(r.command.Users, ", "),
	}
}
func (r RosterPayload)toJson()[]byte{
	b, _ := json.Marshal(&r)
	return b
}
func (r *RosterPayload)FromJson(b []byte)error{
	return json.Unmarshal(b, r)
}

func (r *RosterPayload) fromMap(m map[string]string) {
	r.Channel = m["channel"]
	r.StartTime, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", m["starttime"])
	r.Message = m["message"]
	r.Users = strings.Split(m["users"], ", ")
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	s := server{Client: nil}
	s.SubscribeHandler(w, r)
}

/* SubscribeHandler subscribes a slack channel to a recurring message */
func (s *server) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Write([]byte("Processing request"))
	cmd, _ := slack.SlashCommandParse(r)
	_, err := Subscribe(cmd)
	if err != nil {
		log.Println("{" + err.Error() + "}")
	} else {
		log.Println("yay")
	}

}

func Subscribe(cmd slack.SlashCommand) (*pubsub.Subscription, error) {
	rosterbotCommand, err := ParseCommand(cmd.Text)
	if err != nil {
		return nil, err
	}
	payload := RosterPayload{command: rosterbotCommand, Channel: cmd.ChannelID, Token: cmd.Token}
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, "joshcarp-installer")
	if err != nil {
		return nil, err
	}
	return pubsubService.CreateSubscription(ctx, payload.Channel+strconv.Itoa(rand.Int()), pubsub.SubscriptionConfig{
		Topic: pubsubService.Topic("slack"),
		PushConfig: pubsub.PushConfig{
			Endpoint:   os.Getenv("PUSH_URL")+"?content="+base64.StdEncoding.EncodeToString(payload.toJson()),
			Attributes: payload.toMap(),
		},
	})
}

func Unsubscribe(cmd slack.SlashCommand) (*pubsub.Subscription, error) {
	rosterbotCommand, err := ParseCommand(cmd.Text)
	if err != nil {
		return nil, err
	}
	payload := RosterPayload{command: rosterbotCommand, Channel: cmd.ChannelID}
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, "joshcarp-installer")
	if err != nil {
		return nil, err
	}
	return pubsubService.CreateSubscription(ctx, payload.Channel, pubsub.SubscriptionConfig{
		Topic: pubsubService.Topic("slack"),
		PushConfig: pubsub.PushConfig{
			Endpoint:   os.Getenv("PUSH_URL"),
			Attributes: payload.toMap(),
		},
		Filter: filter.CreateFilter(payload.Time),
	})
}
func CreateURL(){

}
func PublishHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, "joshcarp-installer")
	if err != nil {
		return
	}
	res := pubsubService.Topic("slack").Publish(ctx, &pubsub.Message{
		ID:              "foobar123",
		Data:            []byte("{'content': '1234'}"),
		Attributes:      cron.Now().Map(),
		PublishTime:     time.Now(),
	})
	for{
		select {
		case <-res.Ready():
			return
		}
	}
}


func RespondHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := httputil.DumpRequest(r, true)
	contents, _ := base64.StdEncoding.DecodeString(r.URL.Query().Get("content"))
	payload := RosterPayload{}
	payload.FromJson(contents)
	fmt.Println(string(b))
	c := slack.New(payload.Token)
	message := fmt.Sprintf(`{
  "channel": "%s",
  "text": "%s",
  "as_user": true
}`, payload.Channel, payload.Message)
	c.PostMessage(payload.Channel, slack.MsgOptionText(message, false))
}
