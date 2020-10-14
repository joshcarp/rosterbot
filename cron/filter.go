package cron

import (
	"fmt"
	"strings"
)

func p(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func CreateFilter(c Cron) (string) {
	filter := []string{}
	switch c.Minute {
	case "*":
	default:
		filter = append(filter, p(`(attributes.minute = "%s")`, c.Minute))
	}
	switch c.Hour {
	case "*":
	default:
		filter = append(filter, p(`(attributes.hour = "%s")`, c.Hour))
	}
	switch c.Dom {
	case "*":
	default:
		filter= append(filter,  p(`(attributes.dom = "%s")`, c.Dom))
	}
	switch c.Month {
	case "*":
	default:
		filter = append(filter,p(`(attributes.month = "%s")`, c.Month))
	}
	switch c.Dow {
	case "*":
	default:
		filter = append(filter,p(`(attributes.dow = "%s")`, c.Dow))
	}
	return strings.Join(filter, " AND ")
}
