package roster

import (
	"context"
	"github.com/slack-go/slack"
)

func (s Server) CreateSecret(name string, payload interface{}) error {
	ctx := context.Background()
	_, err := s.Firebase.Collection("webhooks").Doc(name).Set(ctx, payload)
	return err
}

func (s Server) GetSecret(name string) (slack.OAuthV2Response, error) {
	doc := s.Firebase.Collection("webhooks").Doc(name)
	b, err := doc.Get(context.Background())
	if err != nil {
		return slack.OAuthV2Response{}, err
	}
	var secret slack.OAuthV2Response
	err = b.DataTo(&secret)
	if err != nil {
		return slack.OAuthV2Response{}, err
	}
	return secret, nil
}
