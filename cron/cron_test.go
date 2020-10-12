package cron

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseCommand(t *testing.T) {
	type testcase struct {
		in      string
		out    string
	}
	var tests = []testcase{
		{in:"* * * * *", out:"* * * * *"},
		{in:"0 3 9 5 7", out:"0 3 9 5 7"},
		{in:" 0)(DMAN9NSDJ98 08audsfjasd8f asdasd monday asdasd", out:" 0)(DMAN9NSDJ98 08audsfjasd8f asdasd monday asdasd"},
	}
	for _, test := range tests{
		t.Run(test.in, func(t *testing.T) {
			actual, err := Parse(test.in)
			require.NoError(t, err)
			require.Equal(t, test.out, actual.String())

		})
	}
}
