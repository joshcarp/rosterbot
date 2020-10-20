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
	payload := command.RosterPayload{Command: rosterbotCommand, ChannelID: cmd.ChannelID, Token: cmd.Token, TeamID: cmd.TeamID}
	webhookDoc := s.Firebase.Collection("webhooks").Doc(payload.TeamID + "-" + payload.ChannelID)
	a, err := webhookDoc.Get(ctx)
	if err != nil {
		return payload, time.Now(), notsubscribederror
	}
	if a.Data() == nil {
		return payload, time.Now(), notsubscribederror
	}
	_, err = s.Firebase.Collection("subscriptions").Doc(payload.ChannelID+strconv.Itoa(rand.Int())).Set(ctx, payload)
	return payload, payload.Time.Next(time.Now()), err
}
