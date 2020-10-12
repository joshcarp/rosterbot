package rosterbot

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCommand(t *testing.T) {
	type testcase struct {
		in      string
		message string
		users   []string
		time    string
	}
	var tests = []testcase{
		{
			in:      `\roster "0 0 9 * *", "this is the message", @user1, @user2 `,
			message: "this is the message",
			users:   []string{"@user1", "@user2"},
			time:    "0 0 9 * *",
		},
		{
			in:      `\roster "0 0 9 * *" "this is the message", @user1,@user2 `,
			message: "this is the message",
			users:   []string{"@user1", "@user2"},
			time:    "0 0 9 * *",
		},
		{
			in:      `\roster "0 0 9 * *" "this is the message", @user1,@user2 `,
			message: "this is the message",
			users:   []string{"@user1", "@user2"},
			time:    "0 0 9 * *",
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			command, err := ParseCommand(test.in)
			require.NoError(t, err)
			require.Equal(t, test.users, command.Users)
			require.Equal(t, test.time, command.Time.String())
			require.Equal(t, test.message, command.Message)
		})
	}

}
