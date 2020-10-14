package roster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/joshcarp/rosterbot/command"
	"github.com/joshcarp/rosterbot/secrets"
	"github.com/slack-go/slack"
)

func (s Server) Respond(ctx context.Context, contents []byte) error {
	payload := command.RosterPayload{}
	if err := payload.FromJson(contents); err != nil {
		return err
	}
	b, err := secrets.GetSecretData(payload.TeamID + "-" + payload.ChannelID)
	if err != nil {
		return fmt.Errorf("Error getting secret data %w", err)
	}
	var secret slack.OAuthV2Response
	if err := json.Unmarshal(b, &secret); err != nil {
		return err
	}
	message := payload.Message
	if len(payload.Users) > 0{
		message += " "+ payload.Users[payload.Time.Steps(payload.StartTime, time.Now())%len(payload.Users)]
	}
	if err := slack.PostWebhookCustomHTTPContext(
		ctx,
		secret.IncomingWebhook.URL,
		s.Client,
		&slack.WebhookMessage{
			Username: secret.BotUserID,
			Text:     payload.Message,
		}); err != nil {
		return err
	}
	return nil
}

func respond(){

}