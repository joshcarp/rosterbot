package roster

import (
	"context"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/slack-go/slack"
)

func (s Server)Unsubscribe(cmd slack.SlashCommand) (int, error) {
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, s.ProjectID)
	if err != nil {
		return 0, err
	}
	subs := pubsubService.Subscriptions(ctx)
	unsubbed := 0
	for {
		sub, _ := subs.Next()
		if sub == nil {
			return unsubbed, nil
		}
		if strings.HasPrefix(sub.ID(), cmd.ChannelID) {
			if err := sub.Delete(ctx); err != nil {
				return unsubbed, err
			}
			unsubbed++
		}
	}
}
