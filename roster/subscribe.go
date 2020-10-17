package roster

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/joshcarp/rosterbot/command"
	"github.com/slack-go/slack"
)

func (s Server) Subscribe(ctx context.Context, cmd slack.SlashCommand) (command.RosterPayload, time.Time, error) {
	rosterbotCommand, err := command.ParseCommand(cmd.Text)
	if err != nil {
		return command.RosterPayload{}, time.Time{}, err
	}
	payload := command.RosterPayload{Command: rosterbotCommand, ChannelID: cmd.ChannelID, Token: cmd.Token, TeamID: cmd.TeamID}
	_, err = s.Firebase.Collection("subscriptions").Doc(payload.ChannelID+strconv.Itoa(rand.Int())).Set(ctx, payload)
	return payload, payload.Time.Next(time.Now()), err
}
