package roster

import (
	"context"
	"github.com/slack-go/slack"
)

func (s Server) Enroll(ctx context.Context, code string) error {
	accessToken, err := slack.GetOAuthV2ResponseContext(
		ctx,
		s.Client,
		s.SlackClientID,
		s.SlackClientSecret,
		code,
		"")
	if err != nil {
		return err
	}
	return s.CreateSecret(accessToken.Team.ID+"-"+accessToken.IncomingWebhook.ChannelID, accessToken)
}
