package rosterbot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
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

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, "joshcarp-installer")
	if err != nil {
		return
	}
	pubsubService.Topic("slack").Publish(ctx, &pubsub.Message{
		ID:              "",
		Data:            nil,
		Attributes:      cron.Now().Map(),
		PublishTime:     time.Time{},
		DeliveryAttempt: nil,
		OrderingKey:     "",
	})
}


func RespondHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := httputil.DumpRequest(r, true)
	fmt.Println(b)
}
