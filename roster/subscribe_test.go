package roster

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/joshcarp/rosterbot/database"
	"github.com/slack-go/slack"
)

type testcase struct {
	text      string
	teamid    string
	channelid string
	message   map[string]string
}

var tests = []testcase{

	{
		text:      `add "* * * * *" "message" @user1`,
		teamid:    "team1",
		channelid: "channel1",
		message: map[string]string{"3:04PM":"message @user1"},
	},
}

func TestFilter(t *testing.T) {
	for _, test := range tests {
		ctx := context.Background()
		cmd := slack.SlashCommand{
			TeamID:    test.teamid,
			ChannelID: test.channelid,
			Text:      test.text,
		}
		s, client := NewMockServer(test)
		_, err := s.Enroll(ctx, "")
		require.NoError(t, err)
		_, _, err = s.Subscribe(ctx, cmd)
		require.NoError(t, err)

		for timestr, message := range test.message{
			t1, err := time.Parse(time.Kitchen, timestr)
			require.NoError(t, err)
			err = s.Respond(ctx, t1)
			require.NoError(t, err)
			require.Equal(t, message, client.message)
		}
	}

}

func NewMockServer(test testcase) (Server, *testSlackClient) {
	testClient := testSlackClient{
		test: test,
	}
	cl := httptest.NewServer(&testClient)
	testClient.Client = cl.Client()
	u, _ := url.Parse(cl.URL)
	testClient.url = u
	s, _ := server(testClient)
	return s, &testClient
}

func server(client HttpClient) (Server, error) {
	a, err := database.NewMap()
	if err != nil {
		return Server{}, err
	}
	return NewServer("", "foobar", "", a, client), nil
}

type testSlackClient struct {
	url *url.URL
	*http.Client
	message slack.WebhookMessage
	test    testcase
}

func (t *testSlackClient) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/oauth.v2.access":
		a, _ := json.Marshal(slack.OAuthV2Response{
			AccessToken: "AccessToken",
			TokenType:   "TokenType",
			Team: slack.OAuthV2ResponseTeam{
				ID: t.test.teamid,
			},
			IncomingWebhook: slack.OAuthResponseIncomingWebhook{
				URL:       t.url.String(),
				ChannelID: t.test.channelid,
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
		t.message = message
	}

}

func (t testSlackClient) Do(req *http.Request) (*http.Response, error) {
	req.URL.Host = t.url.Host
	req.URL.Scheme = t.url.Scheme
	req.Host = t.url.Host
	return t.Client.Do(req)
}
