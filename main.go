package main

import (
	"log"
)

func main() {
	store, err := NewPgStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := newAPIServer(":3000", store)

	server.Run()

}
