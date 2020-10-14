package roster

import (
	"context"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"

	"github.com/joshcarp/rosterbot/cron"

	"cloud.google.com/go/pubsub"
	"github.com/joshcarp/rosterbot/command"
	"github.com/slack-go/slack"
)

func (s Server) Subscribe(ctx context.Context, cmd slack.SlashCommand) (*pubsub.Subscription, time.Time, error) {
	rosterbotCommand, err := command.ParseCommand(cmd.Text)
	if err != nil {
		return nil, time.Time{}, err
	}
	payload := command.RosterPayload{Command: rosterbotCommand, ChannelID: cmd.ChannelID, Token: cmd.Token, TeamID: cmd.TeamID}
	pubsubService, err := pubsub.NewClient(ctx, s.ProjectID)
	if err != nil {
		return nil, time.Time{}, err
	}
	sub, err := pubsubService.CreateSubscription(ctx, payload.ChannelID+strconv.Itoa(rand.Int()), pubsub.SubscriptionConfig{
		Topic:  pubsubService.Topic(s.Topic),
		Filter: cron.CreateFilter(payload.Time),
		PushConfig: pubsub.PushConfig{
			Endpoint:   s.PushURL + "?content=" + base64.StdEncoding.EncodeToString(payload.ToJson()),
			Attributes: payload.ToMap(),
		},
	})
	return sub, payload.Time.Next(time.Now()), err
}
