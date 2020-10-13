package roster

import "net/http"

type Server struct {
	Client            *http.Client
	Topic             string
	PushURL           string
	ProjectID         string
	SlackClientID     string
	SlackClientSecret string
}

func NewServer(topic, pushURL, projectID string) Server {
	return Server{
		Client:    http.DefaultClient,
		Topic:     topic,
		PushURL:   pushURL,
		ProjectID: projectID,
	}
}
