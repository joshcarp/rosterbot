package roster

import (
	"fmt"

	"github.com/slack-go/slack"
)

func (s Server) Unsubscribe(cmd slack.SlashCommand) (string, error) {
	cols, err := s.Database.Filter("subscriptions", "==", "", map[string]interface{}{"ChannelID": cmd.ChannelID, "TeamID": cmd.TeamID})
	if err != nil {
		return "", err
	}
	unsubbed := 0
	for _, sub := range cols {
		if err := s.Database.Delete("subscriptions", sub.ID); err == nil {
			unsubbed++
		}
	}
	return fmt.Sprintf("Number of rosters unsubscribed from: %d", unsubbed), nil
}
