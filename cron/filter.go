package cron

import (
	"fmt"
)

func p(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func CreateFilter(c Cron) (filter string) {
	switch c.Minute {
	case "*":
	default:
		filter += p(`(attributes.minute = "%s")`, c.Minute)
	}
	switch c.Hour {
	case "*":
	default:
		filter += p(`AND (attributes.hour = "%s")`, c.Hour)
	}
	switch c.Dom {
	case "*":
	default:
		filter += p(`AND (attributes.dom = "%s")`, c.Dom)
	}
	switch c.Month {
	case "*":
	default:
		filter += p(`AND (attributes.month = "%s")`, c.Month)
	}
	switch c.Dow {
	case "*":
	default:
		filter += p(`AND (attributes.dow = "%s")`, c.Dow)
	}
	return
}
