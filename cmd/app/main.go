package main

import (
	"github.com/0z0sk0/intersvyaz-go-test/config"
	"github.com/0z0sk0/intersvyaz-go-test/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.CreateApp()

	log.Printf("PORT: %s", viper.GetString("app.port"))

	if err := app.Start(viper.GetString("app.port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
