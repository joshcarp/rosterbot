package roster

import (
	"context"
	"github.com/slack-go/slack"
)

func (s Server) Enroll(ctx context.Context, code string) (*slack.OAuthV2Response, error) {
	accessToken, err := slack.GetOAuthV2ResponseContext(
		ctx,
		s.Client,
		s.SlackClientID,
		s.SlackClientSecret,
		code,
		"")
	if err != nil {
		return accessToken, err
	}
	return accessToken, s.Database.Set("webhooks", accessToken.Team.ID+"-"+accessToken.IncomingWebhook.ChannelID, accessToken)
}
