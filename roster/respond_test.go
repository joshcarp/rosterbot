package roster

import (
	"context"
	"testing"
	"time"
)

func TestRespond(t *testing.T) {
	s, _ := server(nil)
	s.Respond(context.Background(), time.Now())
	//server().Unsubscribe(cmd)
}