package main

import (
	"log"

	_ "github.com/bibin-zoz/api-gateway/docs" // Import the docs package
	"github.com/bibin-zoz/api-gateway/pkg/config"
	"github.com/bibin-zoz/api-gateway/pkg/di"
)

// @title Social-Match
// @version 1.0
// @description SOCIAL MATCH API Gateway server.
// @termsOfService http://swagger.io/terms/

// @host exclusivestore.xyz
// @BasePath /

// @securityDefinitions.apikey UserAuthorization
// @in header
// @name AccessToken
// @securityDefinitions.apikey UserConfirmToken
// @in header
// @name ConfirmToken
// @securityDefinitions.apikey AdminAutherisation
// @in header
// @name AccessToken
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
