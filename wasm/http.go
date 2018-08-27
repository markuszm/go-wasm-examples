package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, err := client.Get("https://dbf.finalrewind.org/FD?mode=marudor&version=3")

	if err != nil {
		log.Fatal("error in http request", err)
	}

	msg, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal("error reading body", err)
	}

	log.Print(string(msg))
}
