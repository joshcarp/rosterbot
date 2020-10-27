package roster

import (
	"net/http"

	"github.com/joshcarp/rosterbot/database"
)

type Server struct {
	Client            HttpClient
	PushURL           string
	SlackClientID     string
	SlackClientSecret string
	Database          database.Database
}

func NewServer(pushURL, slackClientID, slackClientSecret string, db database.Database, client HttpClient) Server {
	return Server{
		Client:            client,
		PushURL:           pushURL,
		SlackClientID:     slackClientID,
		SlackClientSecret: slackClientSecret,
		Database:          db,
	}
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}