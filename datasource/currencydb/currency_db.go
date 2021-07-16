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
	currencyMaster = make(map[string]currencydb.CurrencyMaster)
	config         = configuration.Default()
)

func init() {
	loadValidSymbols(config)
	go InitWebsocket(config)
}

// InitWebsocket init web socket.
func InitWebsocket(c *configuration.Configuration) {
	flag.Parse()
	message := make(chan util.MessageChannel)
	client, err := util.NewWebSocketClient(c.WSScheme, c.APIHost, c.WSAPIPath, message)
	if err != nil {
		panic(err)
	}
	for d := range message {

		m := currencydb.TickerResponse{}
		if err := json.Unmarshal(d.Data, &m); err != nil {
			log.Printf("Unable to parse message.%s", string(d.Data))
		}

		if m.ID == 101 && m.Result.ID != "" {
			sym := &currencydb.CurrencyMaster{}
			switch s := m.Result.ID; s {
			case "ETH":
				sym = GetSymbol("ETHBTC")
			case "BTC":
				sym = GetSymbol("BTCUSD")
			default:
				sym = GetSymbol(m.Result.ID)
			}
			sym.FullName = m.Result.FullName
			currencyMaster[sym.ID] = *sym
		} else if m.ID != 101 {
			sym := GetSymbol(m.Params.Symbol)
			sym.Params.Ask = m.Params.Ask
			sym.Params.Bid = m.Params.Bid
			sym.Params.Last = m.Params.Last
			sym.Params.Open = m.Params.Open
			sym.Params.Low = m.Params.Low
			sym.Params.High = m.Params.High
			currencyMaster[m.Params.Symbol] = *sym
		}
		fmt.Println(m)
	}

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	client.Stop()
}

func loadValidSymbols(c *configuration.Configuration) {
	var url = fmt.Sprintf("%s://%s%s", c.APIScheme, c.APIHost, c.SymbolAPIPath)
	cm := []currencydb.CurrencyMaster{}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cm); err != nil {
		panic(err)
	}
	for _, c := range cm {
		currencyMaster[c.ID] = c
	}
}

// IsValidSymbol Checks if symbol is valid.
func IsValidSymbol(symbol string) (*currencydb.CurrencyMaster, bool) {
	s := GetSymbol(symbol)
	if s != nil {
		return s, true
	}
	return nil, false
}

// GetSymbol gets the sumbol from inmemory database.
func GetSymbol(symbol string) *currencydb.CurrencyMaster {
	s := currencyMaster[symbol]
	fmt.Println(s)
	if s.ID == symbol {
		return &s
	}
	for _, v := range currencyMaster {
		if v.BaseCurrency == symbol {
			return &v
		}
	}
	return nil
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

// GetAllSymbol gets all supported sumbol.
func GetAllSymbol() *[]currencydb.CurrencyMaster {
	result := []currencydb.CurrencyMaster{}
	for _, s := range config.SupportedSymbols {
		result = append(result, currencyMaster[s])
	}
	return &result
}
