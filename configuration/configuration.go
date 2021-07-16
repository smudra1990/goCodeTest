package configuration

import (
	"encoding/json"
	"os"
)

//Configuration struct represents configuration for application.
type Configuration struct {
	HTTPServerPort   int64    `json:"http_server_port,omitempty"`
	SupportedSymbols []string `json:"supported_symbols,omitempty"`
	APIHost          string   `json:"api_host,omitempty"`
	SymbolAPIPath    string   `json:"symbol_api_path,omitempty"`
	APIScheme        string   `json:"api_scheme,omitempty"`
	WSScheme         string   `json:"ws_scheme,omitempty"`
	TickerMethod     string   `json:"ticker_method,omitempty"`
	SymbolMethod     string   `json:"symbol_method,omitempty"`
	WSAPIPath        string   `json:"wsapi_path,omitempty"`
}

//Default configiration
func Default() *Configuration {
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
