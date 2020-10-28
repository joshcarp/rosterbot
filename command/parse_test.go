package command

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
			in:      `\roster add "0 0 9 * *", "this is the message", @user1, @user2 `,
			message: "this is the message",
			users:   []string{"@user1", "@user2"},
			time:    "0 0 9 * *",
		},
		{
			in:      `\roster add "0 0 9 * *" "this is the message", @user1,@user2 `,
			message: "this is the message",
			users:   []string{"@user1", "@user2"},
			time:    "0 0 9 * *",
		},
		{
			in:      `\roster add "0 0 9 * *" "this is the message", @user1,@user2 `,
			message: "this is the message",
			users:   []string{"@user1", "@user2"},
			time:    "0 0 9 * *",
		},
		{
			in:      `/roster add "* 9 * * *" "This should execute at 7pm" @joshuacarpeggiani`,
			message: "This should execute at 7pm",
			time:    "* 9 * * *",
			users:   []string{"@joshuacarpeggiani"},
		},
		{
			in: `/roster add “* * * * *” “message“ 
 <@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> 
{skip} {skip}
<@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> <@U1234|user> <#C1234|general> 
{skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user>
 {skip} {skip}
<@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user>
 {skip} {skip}
<@U1234|user>  <@U1234|user> <@U1234|user> <@U1234|user> <@U1234|user> 
 {skip} {skip}`,
			message: "message",
			time:    "* * * * *",
			users:   []string{"<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "{skip}", "{skip}", "<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "<@U1234|user>", "<#C1234|general>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "<@U1234|user>", "{skip}", "{skip}"},
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

func TestParseUsers(t *testing.T) {

}
