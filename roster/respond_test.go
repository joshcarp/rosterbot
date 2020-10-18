package roster

import (
	"context"
	"testing"
	"time"
)

/*
Day.Monday
Day.Tuesday
Day.Wednesday
Day.Thursday

*/
func TestRespond(t *testing.T) {
	server().Respond(context.Background(), time.Now())
	//server().Unsubscribe(cmd)
}