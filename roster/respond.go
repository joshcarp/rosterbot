package roster

import (
	"context"
	"fmt"
	"github.com/joshcarp/rosterbot/command"
	"time"

	"github.com/joshcarp/rosterbot/cron"

	"github.com/slack-go/slack"
)

func (s Server) Respond(ctx context.Context, time2 time.Time) error {
	c := cron.Time(time2)
	iter := s.Firebase.Collection("subscriptions").Where("Time/Minute", "==", c.Minute).Documents(ctx)
	docs, _ := iter.GetAll()
	for _, doc := range docs{
		var payload command.RosterPayload
		doc.DataTo(&payload)
		webhookDoc := s.Firebase.Collection("webhooks").Doc(payload.TeamID+"-"+payload.ChannelID)
		if webhookDoc == nil{
			return fmt.Errorf("Channel not authorized")
		}
		snap, err := webhookDoc.Get(ctx)
		if err != nil{
			return err
		}
		var webhook slack.OAuthV2Response
		snap.DataTo(&webhook)
		message := payload.Message
		if len(payload.Users) > 0{
			message += " "+ payload.Users[payload.Time.Steps(payload.StartTime, time.Now())%len(payload.Users)]
		}
		go slack.PostWebhookCustomHTTPContext(
			ctx,
			webhook.IncomingWebhook.URL,
			s.Client,
			&slack.WebhookMessage{
				Text:     message,
			})
	}

	return nil
}
