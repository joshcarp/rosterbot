package roster

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
)

type Server struct {
	Client            *http.Client
	Topic             string
	PushURL           string
	ProjectID         string
	SlackClientID     string
	SlackClientSecret string
	Firebase          *firestore.Client
}

func NewServer(topic, pushURL, projectID, slackClientID, slackClientSecret string) Server {
	firebase, _ := firestore.NewClient(context.Background(), projectID)
	return Server{
		Client:            http.DefaultClient,
		Topic:             topic,
		PushURL:           pushURL,
		ProjectID:         projectID,
		SlackClientID:     slackClientID,
		SlackClientSecret: slackClientSecret,
		Firebase:          firebase,
	}
}
