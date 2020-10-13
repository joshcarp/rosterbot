package rosterbot

import (
	"encoding/hex"
	"fmt"
	"github.com/slack-go/slack"
	"math/rand"
	"testing"
)

func TestRosterBot(t *testing.T){
	NewMockRequest(map[string][]string{
		"token":{"1234"},
		"team_id":{"567"},
		"team_domain":{"890"},
		"enterprise_id":{"123"},
		"enterprise_name":{"234"},
		"channel_id":{"acsf"},
		"user_id":{"a sdf"},
		"user_name":{"casfd"},
		"command":{"roster "},
		"text":{"acs"},
		"response_url":{"acf"},
		"trigger_id":{"csf"},
	},
	)

}
func NewMockRequest(m map[string][]string){
	//s := server{}
	//ser := httptest.NewServer(&s)
	//ser.Client().PostForm(ser.URL, m)
}

func TestSlackCommandSubscribe(t *testing.T){
	Subscribe(slack.SlashCommand{
		Token:          "",
		TeamID:         "",
		TeamDomain:     "",
		EnterpriseID:   "",
		EnterpriseName: "",
		ChannelID:      "",
		ChannelName:    "",
		UserID:         "",
		UserName:       "",
		Command:        "\\roster \"0 0 9 * *\", \"this is the message\", @user1, @user2",
		Text:           "",
		ResponseURL:    "",
		TriggerID:      "",
	})
}



func Salt() (string) {
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	hexSalt := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(hexSalt, bytes)
	return string(hexSalt)
}
