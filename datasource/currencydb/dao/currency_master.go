package currencydb

// CurrencyMaster data access object.
type CurrencyMaster struct {
	ID                   string `json:"id,omitempty"`
	FullName             string `json:"fullName,omitempty"`
	BaseCurrency         string `json:"baseCurrency,omitempty"`
	QuoteCurrency        string `json:"quoteCurrency,omitempty"`
	QuantityIncrement    string `json:"quantityIncrement,omitempty"`
	TickSize             string `json:"tickSize,omitempty"`
	TakeLiquidityRate    string `json:"takeLiquidityRate,omitempty"`
	ProvideLiquidityRate string `json:"provideLiquidityRate,omitempty"`
	FeeCurrency          string `json:"feeCurrency,omitempty"`
	Params               struct {
		Ask    string `json:"ask,omitempty"`
		Bid    string `json:"bid,omitempty"`
		Last   string `json:"last,omitempty"`
		Open   string `json:"open,omitempty"`
		Low    string `json:"low,omitempty"`
		High   string `json:"high,omitempty"`
		Symbol string `json:"symbol,omitempty"`
	} `json:"params,omitempty"`
}

//TickerResponse message from subscribed ticker.
type TickerResponse struct {
	ID     int    `json:"id,omitempty"`
	Method string `json:"method,omitempty"`
	Params struct {
		Ask    string `json:"ask,omitempty"`
		Bid    string `json:"bid,omitempty"`
		Last   string `json:"last,omitempty"`
		Open   string `json:"open,omitempty"`
		Low    string `json:"low,omitempty"`
		High   string `json:"high,omitempty"`
		Symbol string `json:"symbol,omitempty"`
	} `json:"params,omitempty"`
	Result struct {
		ID       string `json:"id,omitempty"`
		FullName string `json:"fullName,omitempty"`
	} `json:"result,omitempty"`
}
