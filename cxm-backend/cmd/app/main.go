package main

import (
	"log"

	"cxm-auth/config"
	"cxm-auth/server"

	"github.com/spf13/viper"
)

func main() {
	config.Init()

	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := app.Run(viper.GetString("PORT")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
