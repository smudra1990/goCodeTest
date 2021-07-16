package app

import (
	"github.com/smudra1990/goCodeTest/configuration"
	"github.com/smudra1990/goCodeTest/controllers"
	"github.com/smudra1990/goCodeTest/vodka"
)

var (
	router = vodka.Default()
	config = &configuration.Configuration{}
)

func init() {
	config = configuration.Default()
}

//StartApplication will prepare all configuration and start http server
func StartApplication() {
	mapUrls()
	router.Run(config.HTTPServerPort)
}

func mapUrls() {
	router.GET("/currency/all", controllers.GetAll)
	router.GET("/currency/{symbol}", controllers.Get)
}
