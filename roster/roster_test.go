package roster

import (
	"context"
	"testing"
	"time"
)

func TestRoster(t *testing.T) {
	a, _ := server(nil)
	a.Respond(context.Background(), time.Now())
}