package rosterbot

import (
	"fmt"
	"github.com/joshcarp/rosterbot/cron"
	"regexp"
	"strings"
	"time"
)

type command struct{
	StartTime time.Time
	Time    cron.Cron
	Message string
	Users   []string
}

func ParseCommand(cmd string)(command, error){
	var (
		ret = command{Users: []string{}}
		commandRe  = regexp.MustCompile(`\s+?"(?P<time>.*?)"\s*,?\s*"(?P<message>.*?)",?\s*(?P<users>@.+)`)
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
					if err != nil{
						return command{}, err
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
		return ret, fmt.Errorf("Invalid command: %s", cmd)
	}
	return ret, nil
}

func ParseUsers(s string)[]string{
	var ret = []string{}
	withoutcommas := strings.ReplaceAll(s, " ", ",")
	for _, user := range strings.Split(withoutcommas, ","){
		if user != ""{
			ret = append(ret, user)
		}
	}
	return ret
}