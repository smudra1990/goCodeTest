package datasource

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/smudra1990/goCodeTest/configuration"
	currencydb "github.com/smudra1990/goCodeTest/datasource/currencydb/dao"
	"github.com/smudra1990/goCodeTest/util"
)

var (
	currencyMaster = []currencydb.CurrencyMaster{}
	config         = configuration.Default()
)

func init() {
	loadValidSymbols(config)
	go InitWebsocket(config)
}

// InitWebsocket init web socket.
func InitWebsocket(c *configuration.Configuration) {
	flag.Parse()
	data := `{
				"method": "subscribeTicker",
				"params": {
		  			"symbol": "ETHBTC"
				},
				"id": 123
	  		}`

	client, err := util.NewWebSocketClient(c.WSScheme, c.APIHost, c.WSAPIPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting")

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)
	client.Write(data)
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	client.Stop()
}

func loadValidSymbols(c *configuration.Configuration) {
	var url = fmt.Sprintf("%s://%s%s", c.APIScheme, c.APIHost, c.SymbolAPIPath)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &currencyMaster); err != nil {
		panic(err)
	}
}

// IsValidSymbol Checks if symbol is valid.
func IsValidSymbol(symbol string) bool {
	for _, v := range currencyMaster {
		if v.ID == symbol {
			return true
		}
	}
	return false
}

// IsValidSymbolSupported Checks if symbol is supported.
func IsValidSymbolSupported(symbol string) bool {
	for _, v := range config.SupportedSymbols {
		if v == symbol {
			return true
		}
	}
	return false
}
