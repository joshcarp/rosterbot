package roster

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
)

type Server struct {
	Client            *http.Client
	PushURL           string
	ProjectID         string
	SlackClientID     string
	SlackClientSecret string
	Firebase          *firestore.Client
}

func NewServer(pushURL, projectID, slackClientID, slackClientSecret string) Server {
	firebase, err := firestore.NewClient(context.Background(), projectID)
	if err != nil{
		panic(err)
	}
	return Server{
		Client:            http.DefaultClient,
		PushURL:           pushURL,
		ProjectID:         projectID,
		SlackClientID:     slackClientID,
		SlackClientSecret: slackClientSecret,
		Firebase:          firebase,
	}
}
