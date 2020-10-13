package roster

import (
	"context"
	"encoding/json"

	"github.com/joshcarp/rosterbot/secrets"
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
	a, err := json.Marshal(accessToken)
	if err != nil {
		return err
	}
	return secrets.CreateSecret(accessToken.Team.ID+"/"+accessToken.IncomingWebhook.ChannelID, a)
}
