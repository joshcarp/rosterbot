package command

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/joshcarp/rosterbot/cron"
)

type Command struct {
	Time      cron.Cron
	Message   string
	Users     []string
}

func MainCommand(s string) string {
	a := strings.Split(s, " ")
	if len(a) > 0 {
		return a[0]
	}
	return ""
}

func ParseCommand(cmd string) (Command, error) {
	var (
		ret       = Command{Users: []string{}}
		commandRe = regexp.MustCompile(`add ("|“|”)(?P<time>.*?)("|“|”)\s*,?\s*("|“)(?P<message>.*?)("|“|”),?\s*(?P<users>(?s).*)`)
		matched = false
	)
	for _, match := range commandRe.FindAllStringSubmatch(cmd, -1) {
		if match == nil {
			continue
		}
		matched = true
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
	if !matched{
		return ret, fmt.Errorf("Error parsing command")
	}
	return ret, nil
}

func ParseUsers(s string) []string {
	var ret = []string{}
	withoutcommas := strings.ReplaceAll(s, " ", ",")
	withoutcommas = strings.ReplaceAll(withoutcommas, "\n", ",")
	for _, user := range strings.Split(withoutcommas, ",") {
		if user != "" {
			ret = append(ret, user)
		}
	}
	return ret
}
