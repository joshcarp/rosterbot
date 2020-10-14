package cron

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseCommand(t *testing.T) {
	type testcase struct {
		in  string
		out string
	}
	var tests = []testcase{
		{in: "* * * * *", out: "* * * * *"},
		{in: "* 9 * * *", out: "* 9 * * *"},
		{in: "0 3 9 5 7", out: "0 3 9 5 7"},
		{in: " 0)(DMAN9NSDJ98 08audsfjasd8f asdasd monday asdasd", out: " 0)(DMAN9NSDJ98 08audsfjasd8f asdasd monday asdasd"},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			actual, err := Parse(test.in)
			require.NoError(t, err)
			require.Equal(t, test.out, actual.String())

		})
	}
}
func TestWhatever(t *testing.T) {
	fmt.Println(Now().Map())
}

func TestCron(t *testing.T){
	c, _ := Parse("* * * * *")
	a := c.Steps(time.Now(), time.Now().Add(time.Hour))
	require.Equal(t, 60, a)

	c, _ = Parse("0 * * * *")
	a = c.Steps(time.Now(), time.Now().Add(time.Hour*24))
	require.Equal(t, 24, a)
}
