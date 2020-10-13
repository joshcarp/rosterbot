package main

import (
	"net/http"

	"github.com/joshcarp/rosterbot"
)

func main() {
	http.HandleFunc("/", rosterbot.RespondHandler)
	http.ListenAndServe(":8081", nil)
}
