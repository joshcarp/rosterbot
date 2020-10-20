package rosterbot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joshcarp/rosterbot/command"

	"github.com/joshcarp/rosterbot/roster"
	"github.com/slack-go/slack"
)

func RespondHandler(w http.ResponseWriter, r *http.Request) {
	if err := server().Respond(context.Background(), time.Now()); err != nil {
		log.Println(err)
	}
}

func Enroll(w http.ResponseWriter, r *http.Request) {
	auth, err := server().Enroll(context.Background(), r.URL.Query().Get("code"));
	if  err != nil {
		log.Println(err)
	}
	w.Write([]byte("Rosterbot installed on "+auth.Team.Name))
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	cmd, _ := slack.SlashCommandParse(r)
	switch strings.ToLower(command.MainCommand(cmd.Text)) {
	case "add":
		_, time, err := server().Subscribe(context.Background(), cmd)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("New roster added: `%s` starting on %s", cmd.Text, time.String())))
	case "remove":
		i, err := server().Unsubscribe(cmd)
		w.Write([]byte(fmt.Sprintf("Unsubscribed %d roster(s)", i) + cmd.Text))
		if err != nil {
			log.Println(err)
			w.Write([]byte("There was a problem unsubscribing"))
		}
	default:
		w.Write([]byte("Command not known, please specify /roster add or /roster remove"))
	}
}

func server() roster.Server {
	return roster.NewServer(os.Getenv("PUSH_URL"), os.Getenv("PROJECT_ID"), os.Getenv("SLACK_CLIENT_ID"), os.Getenv("SLACK_CLIENT_SECRET"))
}
