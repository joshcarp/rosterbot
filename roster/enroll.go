package roster

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
)

func (s Server) Enroll(ctx context.Context, code string) (string, error) {
	accessToken, err := slack.GetOAuthV2ResponseContext(
		ctx,
		s.Client,
		s.SlackClientID,
		s.SlackClientSecret,
		code,
		"")
	if err != nil {
		return "", err
	}
	err = s.Database.Set("webhooks", accessToken.Team.ID+"-"+accessToken.IncomingWebhook.ChannelID, accessToken)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Rosterbot Installed on: %s", accessToken.Team.Name), nil
}
