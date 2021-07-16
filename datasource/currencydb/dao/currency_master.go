package currencydb

// CurrencyMaster data access object.
type CurrencyMaster struct {
	ID                   string `json:"id,omitempty"`
	BaseCurrency         string `json:"baseCurrency,omitempty"`
	QuoteCurrency        string `json:"quoteCurrency,omitempty"`
	QuantityIncrement    string `json:"quantityIncrement,omitempty"`
	TickSize             string `json:"tickSize,omitempty"`
	TakeLiquidityRate    string `json:"takeLiquidityRate,omitempty"`
	ProvideLiquidityRate string `json:"provideLiquidityRate,omitempty"`
	FeeCurrency          string `json:"USD,omitempty"`
}
