package roster

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joshcarp/rosterbot/command"
	"github.com/joshcarp/rosterbot/cron"
	"github.com/slack-go/slack"
)

func Unsubscribe(cmd slack.SlashCommand) (*pubsub.Subscription, error) {
	rosterbotCommand, err := command.ParseCommand(cmd.Text)
	if err != nil {
		return nil, err
	}
	payload := command.RosterPayload{Command: rosterbotCommand, ChannelID: cmd.ChannelID}
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, "joshcarp-installer")
	if err != nil {
		return nil, err
	}
	return pubsubService.CreateSubscription(ctx, payload.ChannelID, pubsub.SubscriptionConfig{
		Topic: pubsubService.Topic("slack"),
		PushConfig: pubsub.PushConfig{
			Endpoint:   os.Getenv("PUSH_URL"),
			Attributes: payload.ToMap(),
		},
		Filter: cron.CreateFilter(payload.Time),
	})
}
