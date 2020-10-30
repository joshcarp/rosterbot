package roster

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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
	starttime string
	message   map[string]string
}

var tests = []testcase{
	{
		text:      `add "* * * * *" "message" @user1`,
		teamid:    "team1",
		channelid: "channel1",
		starttime: "3:04PM",
		message:   map[string]string{"3:04PM": "message @user1"},
	},
	{
		text:      `add "* * * * *" "message" @user1 @user2 @user3`,
		teamid:    "team1",
		channelid: "channel1",
		starttime: "3:03PM",
		message: map[string]string{
			"3:04PM": "message @user1",
			"3:05PM": "message @user2",
			"3:06PM": "message @user3",
		},
	},
	{
		text:      `add "0 * * * *" "message" @user1 @user2 @user3`,
		teamid:    "team1",
		channelid: "channel1",
		starttime: "2:30PM",
		message: map[string]string{
			"3:00PM": "message @user1",
			"4:00PM": "message @user2",
			"5:00PM": "message @user3",
		},
	},
	{
		text: `add "0 * * * *" "This is a long complex message" <@id1|user1> <@id1|user1> <@id1|user1> <@id1|user1> <@id1|user1>
{skip} {skip} 
<@id2|user2> <@id2|user2> <@id2|user2> <@id2|user2> <@id2|user2>
{skip} {skip}
<@id3|user3> <@id3|user3> <@id3|user3> <@id3|user3> <@id3|user3>`,
		teamid:    "team1",
		channelid: "channel1",
		starttime: "12:01AM",
		message: map[string]string{
			"1:00AM":  "This is a long complex message <@id1|user1>",
			"2:00AM":  "This is a long complex message <@id1|user1>",
			"3:00AM":  "This is a long complex message <@id1|user1>",
			"4:00AM":  "This is a long complex message <@id1|user1>",
			"5:00AM":  "This is a long complex message <@id1|user1>",
			"8:00AM":  "This is a long complex message <@id2|user2>",
			"9:00AM":  "This is a long complex message <@id2|user2>",
			"10:00AM": "This is a long complex message <@id2|user2>",
			"11:00AM": "This is a long complex message <@id2|user2>",
			"12:00PM": "This is a long complex message <@id2|user2>",
			"3:00PM":  "This is a long complex message <@id3|user3>",
			"4:00PM":  "This is a long complex message <@id3|user3>",
			"5:00PM":  "This is a long complex message <@id3|user3>",
			"6:00PM":  "This is a long complex message <@id3|user3>",
			"7:00PM":  "This is a long complex message <@id3|user3>",
		},
	},
	{
		text:      `add "0 * * * *" "This is a long complex message" <@id1|user1>,<@id1|user1>,<@id1|user1>,<@id1|user1>,<@id1|user1>,{skip},{skip},<@id2|user2>,<@id2|user2>,<@id2|user2>,<@id2|user2>,<@id2|user2>,{skip},{skip},<@id3|user3>, <@id3|user3>,<@id3|user3>,<@id3|user3>,<@id3|user3>`,
		teamid:    "team1",
		channelid: "channel1",
		starttime: "12:01AM",
		message: map[string]string{
			"1:00AM":  "This is a long complex message <@id1|user1>",
			"2:00AM":  "This is a long complex message <@id1|user1>",
			"3:00AM":  "This is a long complex message <@id1|user1>",
			"4:00AM":  "This is a long complex message <@id1|user1>",
			"5:00AM":  "This is a long complex message <@id1|user1>",
			"8:00AM":  "This is a long complex message <@id2|user2>",
			"9:00AM":  "This is a long complex message <@id2|user2>",
			"10:00AM": "This is a long complex message <@id2|user2>",
			"11:00AM": "This is a long complex message <@id2|user2>",
			"12:00PM": "This is a long complex message <@id2|user2>",
			"3:00PM":  "This is a long complex message <@id3|user3>",
			"4:00PM":  "This is a long complex message <@id3|user3>",
			"5:00PM":  "This is a long complex message <@id3|user3>",
			"6:00PM":  "This is a long complex message <@id3|user3>",
			"7:00PM":  "This is a long complex message <@id3|user3>",
		},
	},
}

func TestFilter(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := context.Background()
			cmd := slack.SlashCommand{
				TeamID:    test.teamid,
				ChannelID: test.channelid,
				Text:      test.text,
			}
			s, client := NewMockServer(test)
			_, err := s.Enroll(ctx, "")
			require.NoError(t, err)
			t1, err := time.Parse(time.Kitchen, test.starttime)
			_, err = s.Subscribe(ctx, cmd, t1)
			require.NoError(t, err)
			for timestr, message := range test.message {
				t1, err := time.Parse(time.Kitchen, timestr)
				require.NoError(t, err)
				err = s.Respond(ctx, t1)
				require.NoError(t, err)
				require.Equal(t, message, client.message.Text)
			}
		})
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
	return NewServer("", "", a, client), nil
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
