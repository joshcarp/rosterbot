package command

import (
	"encoding/json"
	"strings"
	"time"
)

type RosterPayload struct {
	Command
	ChannelID string
	Token     string
	TeamID    string
}

func (r RosterPayload) ToMap() map[string]string {
	return map[string]string{
		"channel":   r.ChannelID,
		"starttime": r.Command.StartTime.String(),
		"message":   r.Command.Message,
		"users":     strings.Join(r.Command.Users, ", "),
	}
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