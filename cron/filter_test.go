package cron

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	a, _ := Parse("* 9 * 9 *")
	fmt.Println(CreateFilter(a))
}

func TestFilter2(t *testing.T) {
	a, _ := Parse("0 9 * * *")
	fmt.Println(CreateFilter(a))
}
