package roster

import (
	"context"
	"fmt"
	"github.com/joshcarp/rosterbot/command"
	"github.com/joshcarp/rosterbot/cron"
	"time"

	"github.com/slack-go/slack"
)
/*
Day.Monday
Day.Tuesday
Day.Wednesday
Day.Thursday

*/
func (s Server) Respond(ctx context.Context, time2 time.Time) error {
	filters := cron.Expand(cron.Time(time2))
	col := s.Firebase.Collection("subscriptions")
	q := col.Query
	for filter, val := range filters{
		q = q.Where("Time.Complete."+filter, "==", val)
	}
	iter := q.Documents(ctx)
	if iter == nil{
		return nil
	}
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
