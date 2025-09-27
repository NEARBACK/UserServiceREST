package main

import (
	"log"
	"useservice/internal/entry"
)

func main() {
	server, err := entry.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	server.Start()

	err = server.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
