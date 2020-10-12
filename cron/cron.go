package cron

import (
	"fmt"
	"regexp"
)

var cronRe = regexp.MustCompile(`(?P<minute>.*?) (?P<hour>.*?) (?P<dom>.*?) (?P<month>.*?) (?P<dow>.*)`)

type Cron struct{
	Minute string
	Hour   string
	Dom    string
	Month  string
	Dow    string
}

func (c Cron)String()string{
	return fmt.Sprintf("%s %s %s %s %s", c.Minute, c.Hour, c.Dom, c.Month, c.Dow)
}

func Parse(s string)(Cron, error){
	var ret Cron
	matches := cronRe.FindAllStringSubmatch(s, -1)
	if len(matches) == 0{
		return Cron{}, fmt.Errorf("Can't parse cron")
	}
	for _, match := range matches  {
		if match == nil {
			continue
		}
		for i, name := range cronRe.SubexpNames() {
			if match[i] != "" {
				switch name {
				case "minute":
					ret.Minute = match[i]
				case "hour":
					ret.Hour = match[i]
				case "dom":
					ret.Dom = match[i]
				case "month":
					ret.Month = match[i]
				case "dow":
					ret.Dow = match[i]
				}
			}
		}
	}
	return ret, nil
}