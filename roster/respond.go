package roster

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/joshcarp/rosterbot/cron"

	"github.com/slack-go/slack"
)

/*
Day.Monday
Day.Tuesday
Day.Wednesday
Day.Thursday

*/
func (s Server) Respond(ctx context.Context, time2 time.Time) error {
	var wg sync.WaitGroup
	filters := cron.Expand(cron.Time(time2))
	a, err := s.Database.Filter("subscriptions", "==", "Time.Complete.", filters)
	if err != nil {
		return err
	}
	for _, sub := range a {
		webhook, err := s.GetSecret(sub.TeamID + "-" + sub.ChannelID)
		if err != nil {
			continue
		}
		message := sub.Message
		if len(sub.Users) > 0 {
			steps := sub.Time.Steps(sub.StartTime, time2) - 1
			if steps < 0 {
				steps++
			}
			if sub.Users[steps%len(sub.Users)] == "{skip}" {
				continue
			}
			message += " " + sub.Users[steps%len(sub.Users)]
		}
		wg.Add(1)
		go func() {
			PostWebhookCustomHTTPContext(
				ctx,
				webhook.IncomingWebhook.URL,
				s.Client,
				&slack.WebhookMessage{
					Text: message,
				})
			wg.Done()
		}()

	}
	wg.Wait()
	return nil
}

func PostWebhookCustomHTTPContext(ctx context.Context, url string, httpClient HttpClient, msg *slack.WebhookMessage) error {
	raw, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return errors.Wrap(err, "failed new request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to post webhook")
	}
	defer resp.Body.Close()

	return checkStatusCode(resp)
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode == http.StatusTooManyRequests {
		retry, err := strconv.ParseInt(resp.Header.Get("Retry-After"), 10, 64)
		if err != nil {
			return err
		}
		return &slack.RateLimitedError{time.Duration(retry) * time.Second}
	}

	// Slack seems to send an HTML body along with 5xx error codes. Don't parse it.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}

	return nil
}
