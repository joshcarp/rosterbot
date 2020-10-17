package roster

import (
	"context"
	"strings"

	"github.com/slack-go/slack"
)

func (s Server)Unsubscribe(cmd slack.SlashCommand) (int, error) {
	cols := s.Firebase.Collection("subscriptions").
		Where("ChannelID", "==", cmd.ChannelID).
		Where("TeamID", "==", cmd.TeamID).
		Documents(context.Background())
	unsubbed := 0
	for {
		sub, _ := cols.Next()
		if sub == nil {
			return unsubbed, nil
		}
		if strings.HasPrefix(sub.Data()["ChannelID"].(string), cmd.ChannelID) {
			if _, err := sub.Ref.Delete(context.Background()); err != nil{
				return unsubbed, nil
			}
			unsubbed++
		}
	}
}
