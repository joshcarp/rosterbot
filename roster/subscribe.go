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

var notsubscribederror = fmt.Errorf("error subscribing: This channel is not authorized, to install visit https://slack.com/oauth/v2/authorize?client_id=1367393582980.1445120201280&scope=commands,incoming-webhook&user_scope=")

func (s Server) Subscribe(ctx context.Context, cmd slack.SlashCommand) (command.RosterPayload, time.Time, error) {
	rosterbotCommand, err := command.ParseCommand(cmd.Text)
	if err != nil {
		return command.RosterPayload{}, time.Time{}, err
	}
	payload := command.RosterPayload{
		ID: cmd.ChannelID+strconv.Itoa(rand.Int()),
		Command: rosterbotCommand,
		ChannelID: cmd.ChannelID,
		Token: cmd.Token,
		TeamID: cmd.TeamID}
	_, err = s.GetSecret(payload.TeamID + "-" + payload.ChannelID)
	if err != nil {
		return payload, time.Now(), notsubscribederror
	}
	err = s.Database.Set("subscription", payload.ID, payload)
	return payload, payload.Time.Next(time.Now()), err
}
