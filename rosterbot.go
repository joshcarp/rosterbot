package rosterbot

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joshcarp/rosterbot/database"

	"github.com/joshcarp/rosterbot/command"

	"github.com/joshcarp/rosterbot/roster"
	"github.com/slack-go/slack"
)

func RespondHandler(w http.ResponseWriter, r *http.Request) {
	ser, err := server()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := ser.Respond(context.Background(), time.Now()); err != nil {
		log.Println(err)
	}
}

func Enroll(w http.ResponseWriter, r *http.Request) {
	ser, err := server()
	message, err := ser.Enroll(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		log.Println(err)
		return
	}
	w.Write([]byte(message))
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	ser, err := server()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cmd, _ := slack.SlashCommandParse(r)
	switch strings.ToLower(command.MainCommand(cmd.Text)) {
	case "add":
		message, err := ser.Subscribe(context.Background(), cmd, time.Now())
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(message))
	case "remove":
		message, err := ser.Unsubscribe(cmd)
		w.Write([]byte(message))
		if err != nil {
			log.Println(err)
			w.Write([]byte("There was a problem unsubscribing"))
		}
	default:
		w.Write([]byte("Command not known, please specify /roster add or /roster remove"))
	}
}

func server() (roster.Server, error) {
	fire, err := database.NewFirestore(os.Getenv("PROJECT_ID"))
	if err != nil {
		return roster.Server{}, err
	}
	return roster.NewServer(
		os.Getenv("SLACK_CLIENT_ID"),
		os.Getenv("SLACK_CLIENT_SECRET"),
		fire,
		http.DefaultClient,
	), nil
}
