package main

import (
	"github.com/joshcarp/rosterbot"
	"net/http"
)

func main() {
	http.HandleFunc("/", rosterbot.RespondHandler)
	http.ListenAndServe(":8081", nil)
}
