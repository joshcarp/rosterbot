package command

import (
	"encoding/json"
	"strings"
	"time"
)

type RosterPayload struct {
	ID string
	Command
	ChannelID string
	Token     string
	TeamID    string
	StartTime time.Time
}

func (r RosterPayload) ToJson() []byte {
	b, _ := json.Marshal(&r)
	return b
}

func (r *RosterPayload) FromJson(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *RosterPayload) FromMap(m map[string]string) {
	r.ChannelID = m["channel"]
	r.StartTime, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", m["starttime"])
	r.Message = m["message"]
	r.Users = strings.Split(m["users"], ", ")
}
