package roster

import (
	"context"
	"encoding/base64"
	"github.com/joshcarp/rosterbot/cron"
	"math/rand"
	"strconv"

	"cloud.google.com/go/pubsub"
	"github.com/joshcarp/rosterbot/command"
	"github.com/slack-go/slack"
)

func (s Server) Subscribe(ctx context.Context, cmd slack.SlashCommand) (*pubsub.Subscription, error) {
	rosterbotCommand, err := command.ParseCommand(cmd.Text)
	if err != nil {
		return nil, err
	}
	payload := command.RosterPayload{Command: rosterbotCommand, ChannelID: cmd.ChannelID, Token: cmd.Token, TeamID: cmd.TeamID}
	pubsubService, err := pubsub.NewClient(ctx, s.ProjectID)
	if err != nil {
		return nil, err
	}
	return pubsubService.CreateSubscription(ctx, payload.ChannelID+strconv.Itoa(rand.Int()), pubsub.SubscriptionConfig{
		Topic: pubsubService.Topic(s.Topic),
		Filter: cron.CreateFilter(payload.Time),
		PushConfig: pubsub.PushConfig{
			Endpoint:   s.PushURL + "?content=" + base64.StdEncoding.EncodeToString(payload.ToJson()),
			Attributes: payload.ToMap(),
		},
	})
}
