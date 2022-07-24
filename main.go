package main

import (
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/api"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/collector"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/config"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/environment"
	"github.com/gin-gonic/gin"
)

var configuration config.ConfigReader
var env environment.EnvironmentReader
var imageCollector collector.ImageCollector
var imgCollectorSemaphore chan struct{}

func main() {
	env = environment.LoadEnv()
	configuration = config.NewConfig(env.Port(), env.ApiKey(), env.ConcurrentRequests())
	imgCollectorSemaphore = make(chan struct{}, configuration.ConcurrentRequests())
	imageCollector = collector.NewNasaImageCollector(imgCollectorSemaphore, configuration.ApiKey())

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/pictures", func(ctx *gin.Context) {
		api.GetPictures(ctx, configuration, imageCollector)
	})

	router.Run(configuration.Port())
}
