package roster

import (
	"github.com/slack-go/slack"
)

func (s Server) GetSecret(name string) (slack.OAuthV2Response, error) {
	var secret slack.OAuthV2Response
	err := s.Database.Get("webhooks", name, &secret)
	return secret, err
}
