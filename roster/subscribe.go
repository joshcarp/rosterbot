package roster

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/joshcarp/rosterbot/command"
	"github.com/slack-go/slack"
)

var notsubscribederror = fmt.Errorf("This channel is not authorized, to install visit https://slack.com/oauth/v2/authorize?client_id=1367393582980.1445120201280&scope=commands,incoming-webhook&user_scope=")

func (s Server) Subscribe(ctx context.Context, cmd slack.SlashCommand, starttime time.Time) (string, error) {
	rosterbotCommand, err := command.ParseCommand(cmd.Text)
	if err != nil {
		return "", err
	}
	payload := command.RosterPayload{
		ID:        cmd.ChannelID + strconv.Itoa(rand.Int()),
		Command:   rosterbotCommand,
		ChannelID: cmd.ChannelID,
		Token:     cmd.Token,
		TeamID:    cmd.TeamID,
		StartTime: starttime,
	}
	_, err = s.GetSecret(payload.TeamID + "-" + payload.ChannelID)
	if err != nil {
		return "", notsubscribederror
	}
	err = s.Database.Set("subscriptions", payload.ID, payload)
	return fmt.Sprintf("Roster added, first execution: %v", payload.Time.Next(starttime)), err
}
