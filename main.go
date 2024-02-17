package main

import (
	"context"
	"fmt"
	"go-prompt/config"
	"go-prompt/controller"
	"go-prompt/db"
	"go-prompt/service"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
		log.Warn("DEBUG is true")
	}
	config.Load()

	log.Debugf("Config: %+v", config.GlobalConfig)

	pool, err := db.Connect(context.Background())
	if err != nil {
		log.WithError(err).Fatal("Error connecting to database")
	}
	defer pool.Close()

	q := db.New(pool)
	service := service.NewService(q)

	api := controller.SetupAPI(service)
	log.Panic(api.Run(fmt.Sprintf(":%d", 8080)))
}
