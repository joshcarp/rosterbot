package roster

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/joshcarp/rosterbot/cron"
)

func (s Server) Publish() error {
	ctx := context.Background()
	pubsubService, err := pubsub.NewClient(ctx, s.ProjectID)
	if err != nil {
		return err
	}
	now := cron.Now()
	res := pubsubService.Topic(s.Topic).Publish(
		ctx,
		&pubsub.Message{
			ID:          now.String(),
			Attributes:  now.Map(),
			PublishTime: time.Now(),
		})
	for {
		select {
		case <-res.Ready():
			return nil
		}
	}
}
