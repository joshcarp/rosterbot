package rosterbot

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joshcarp/rosterbot/roster"
	"github.com/slack-go/slack"
)

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	if err := server().Publish(); err != nil {
		log.Println(err)
	}
}

func RespondHandler(w http.ResponseWriter, r *http.Request) {
	contents, _ := base64.StdEncoding.DecodeString(r.URL.Query().Get("content"))
	if err := server().Respond(context.Background(), contents); err != nil {
		log.Println(err)
	}
}

func Enroll(w http.ResponseWriter, r *http.Request) {
	if err := server().Enroll(context.Background(), r.URL.Query().Get("code")); err != nil {
		log.Println(err)
	}
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	cmd, _ := slack.SlashCommandParse(r)
	switch strings.ToLower(cmd.Command){
	case "add":
		if _, err := server().Subscribe(context.Background(), cmd); err != nil {
			log.Println(err)
			w.Write([]byte("Error adding roster"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte("New roster added" +cmd.Text))
	case "remove":
		i, err := server().Unsubscribe(cmd)
		w.Write([]byte(fmt.Sprintf("Unsubscribed %d roster(s)", i) +cmd.Text))
		if err != nil{
			w.Write([]byte("There was a problem unsubscribing"))
		}

	}
}

func server() roster.Server {
	return roster.NewServer(os.Getenv("GCP_TOPIC"), os.Getenv("PUSH_URL"), os.Getenv("PROJECT_ID"), os.Getenv("SLACK_CLIENT_ID"), os.Getenv("SLACK_CLIENT_SECRET"))
}
