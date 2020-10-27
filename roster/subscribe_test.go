package roster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joshcarp/rosterbot/database"
	"github.com/slack-go/slack"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

type test struct{
	text string
	teamid string
	channelid string
}
var tests []test{

}

func TestFilter(t *testing.T) {
	cmd := slack.SlashCommand{
		TeamID:         "id",
		ChannelID:      "foobar",
		Text:           `add "* * * * *" "THis is a message" @joshuacarpeggiani `,
	}
	s, client := NewMockServer()
	s.Enroll(context.Background(), "1234")
	a, b,  c := s.Subscribe(context.Background(), cmd)
	s.Respond(context.Background(), time.Now())
	fmt.Println(a, b, c)
	fmt.Println(client)
}

func NewMockServer()(Server, *testSlackClient){
	testClient := testSlackClient{}
	cl := httptest.NewServer(&testClient)
	testClient.Client = cl.Client()
	u, _ := url.Parse(cl.URL)
	testClient.url = u
	s, _ := server(testClient)
	return s, &testClient
}

func server(client HttpClient) (Server, error) {
	a, err := database.NewMap()
	if err != nil{
		return Server{}, err
	}
	return NewServer("", "foobar", "", a, client), nil
}

type testSlackClient struct{
	url *url.URL
	*http.Client
	messages []slack.WebhookMessage
}

func (t *testSlackClient)ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/api/oauth.v2.access":
		a, _ := json.Marshal(slack.OAuthV2Response{
			AccessToken: "1234",
			TokenType:   "bot",
			Team:        slack.OAuthV2ResponseTeam{
				ID:   "id",
				Name: "name",
			},
			IncomingWebhook: slack.OAuthResponseIncomingWebhook{
				URL:       t.url.String(),
				Channel:   "foobar",
				ChannelID: "foobar",
			},
			Enterprise:    slack.OAuthV2ResponseEnterprise{},
			AuthedUser:    slack.OAuthV2ResponseAuthedUser{},
			SlackResponse: slack.SlackResponse{},
		})
		w.Write(a)
		return
	default:
		message := slack.WebhookMessage{}
		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &message)
		t.messages = append(t.messages, message)
	}

}



func (t testSlackClient)Do(req *http.Request) (*http.Response, error){
	req.URL.Host = t.url.Host
	req.URL.Scheme = t.url.Scheme
	req.Host = t.url.Host
	return t.Client.Do(req)
}