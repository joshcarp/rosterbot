package main

import (
	"github.com/joshcarp/rosterbot"
	"log"
	"net"
	"net/http"

)

func main() {
	http.HandleFunc("/", rosterbot.ServeHTTP)
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(lis, nil))
}
