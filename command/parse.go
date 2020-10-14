package command

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/joshcarp/rosterbot/cron"
)

type Command struct {
	StartTime time.Time
	Time      cron.Cron
	Message   string
	Users     []string
}

func ParseCommand(cmd string) (Command, error) {
	var (
		ret       = Command{Users: []string{}, StartTime: time.Now()}
		commandRe = regexp.MustCompile(`"(?P<time>.*?)"\s*,?\s*"(?P<message>.*?)",?\s*(?P<users>@.+)`)
	)
	for _, match := range commandRe.FindAllStringSubmatch(cmd, -1) {
		if match == nil {
			continue
		}
		for i, name := range commandRe.SubexpNames() {
			if match[i] != "" {
				switch name {
				case "time":
					c, err := cron.Parse(match[i])
					if err != nil {
						return Command{}, err
					}
					ret.Time = c
				case "message":
					ret.Message = match[i]
				case "users":
					ret.Users = ParseUsers(match[i])
				}
			}
		}
	}
	if len(ret.Users) == 0 {
		return ret, fmt.Errorf("Invalid Command: %s", cmd)
	}
	return ret, nil
}

func ParseUsers(s string) []string {
	var ret = []string{}
	withoutcommas := strings.ReplaceAll(s, " ", ",")
	for _, user := range strings.Split(withoutcommas, ",") {
		if user != "" {
			ret = append(ret, user)
		}
	}
	return ret
}
