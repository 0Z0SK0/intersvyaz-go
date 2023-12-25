package main

import (
	"github.com/0z0sk0/intersvyaz-go-test/server"
	"log"
)

func main() {
	app := server.CreateApp()

	if err := app.Start(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
