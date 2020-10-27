package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

var cronRe = regexp.MustCompile(`(?P<minute>.*?) (?P<hour>.*?) (?P<dom>.*?) (?P<month>.*?) (?P<dow>.*)`)

type Cron struct {
	Minute string
	Hour   string
	Dom    string
	Month  string
	Dow    string
	Complete map[string]interface{}
}

func (c Cron) String() string {
	return fmt.Sprintf("%s %s %s %s %s", c.Minute, c.Hour, c.Dom, c.Month, c.Dow)
}

func (c Cron) Map() map[string]string {
	return map[string]string{
		"minute": c.Minute,
		"hour":   c.Hour,
		"dom":    c.Dom,
		"month":  c.Month,
		"dow":    c.Dow,
	}
}

func Parse(s string) (Cron, error) {
	var ret Cron
	matches := cronRe.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return Cron{}, fmt.Errorf("Can't parse cron")
	}
	for _, match := range matches {
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
	ret.Complete = Expand(ret)
	return ret, nil
}

func Expand(c Cron)map[string]interface{}{
	complete := make(map[string]interface{})
	switch c.Minute{
	case "*":
		for i := 0; i < 60; i ++{
			complete["Minute:"+strconv.Itoa(i)] = true
		}
	default:
		complete["Minute:"+c.Minute] = true
	}
	switch c.Hour{
	case "*":
		for i := 0; i < 24; i ++{
			complete["Hour:"+strconv.Itoa(i)] = true
		}
	default:
		complete["Hour:"+c.Hour] = true
	}
	switch c.Month{
	case "*":
		for i := 1; i < 13; i ++{
			complete["Month:"+strconv.Itoa(i)] = true
		}
	default:
		complete["Month:"+c.Month] = true
	}
	switch c.Dom{
	case "*":
		for i := 0; i < 32; i ++{
			complete["Dom:"+strconv.Itoa(i)] = true
		}
	default:
		complete["Dom:"+c.Dom] = true
	}
	switch c.Dow	{
	case "*":
		for i := 1; i < 8; i ++{
			complete["Dow:"+strconv.Itoa(i)] = true
		}
	default:
		complete["Dow:"+c.Dow] = true
	}
	return complete
}
func Now() Cron {
	return Time(time.Now())
}

func Time(t time.Time) Cron {
	t.Hour()
	_, month, day := t.Date()
	return Cron{
		Minute: strconv.Itoa(t.Minute()),
		Hour:   strconv.Itoa(t.Hour()),
		Dom:    strconv.Itoa(day),
		Month:  strconv.Itoa(int(month)),
		Dow:    strconv.Itoa(int(t.Weekday())),
	}
}

func (c Cron) Next(time2 time.Time) time.Time {
	a, _ := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(c.String())
	return a.Next(time2)
}

func (c Cron) Steps(start, end time.Time) int {
	a, _ := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(c.String())
	var steps = 0
	for {
		start = a.Next(start)
		if start.After(end) {
			return steps
		}
		steps++
	}
}
