package main

import (
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/api"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/config"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/environment"
	"github.com/gin-gonic/gin"
)

var configuration config.ConfigReader
var env environment.EnvironmentReader

func main() {
	env = environment.LoadEnv()
	configuration = config.NewConfig(env.Port(), env.ApiKey(), env.ConcurrentRequests())

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/pictures", func(ctx *gin.Context) {
		api.GetPictures(ctx, configuration)
	})

	router.Run(configuration.Port())
}
