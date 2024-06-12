package main

import (
	"log"

	"github.com/bibin-zoz/social-match-room-svc/pkg/config"
	"github.com/bibin-zoz/social-match-room-svc/pkg/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
