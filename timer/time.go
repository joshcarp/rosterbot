package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	for {
		fmt.Println("https://us-central1-joshcarp-installer.cloudfunctions.net/publish")
		http.Get("https://us-central1-joshcarp-installer.cloudfunctions.net/publish")
		time.Sleep(time.Minute)
	}
}
