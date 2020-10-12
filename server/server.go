package main

import (
	 "github.com/joshcarp/whattimeisitrightnow"
	"log"
	"net"
	"net/http"

)

func main() {
	http.HandleFunc("/", whattimeisitrightnow.ServeHTTP)
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(lis, nil))
}
