package app

import (
	"encoding/json"
	"os"

	"github.com/smudra1990/goCodeTest/controllers"
	"github.com/smudra1990/goCodeTest/vodka"
)

var (
	router = vodka.Default()
)

//Configuration struct represents configuration for application.
type Configuration struct {
	HTTPServerPort   int64    `json:"http_server_port,omitempty"`
	SupportedSymbols []string `json:"supported_symbols,omitempty"`
}

//StartApplication will prepare all configuration and start http server
func StartApplication() {
	config := readConfiguration()
	mapUrls()
	router.Run(config.HTTPServerPort)
}

func mapUrls() {
	router.GET("/currency/{symbol}", controllers.GetAll)
	router.GET("/currency/all", controllers.GetAll)
}

func readConfiguration() *Configuration {
	file, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	if err := decoder.Decode(&configuration); err != nil {
		panic(err)
	}
	return &configuration
}
