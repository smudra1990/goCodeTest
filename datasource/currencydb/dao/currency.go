package currencydb

// Currency DAO
type Currency struct {
	ID          string  `json:"id"`
	FullName    string  `json:"full_name"`
	Aslk        float64 `json:"aslk"`
	Bid         float64 `json:"bid"`
	Last        float64 `json:"last"`
	Open        float64 `json:"open"`
	Low         float64 `json:"low"`
	High        float64 `json:"high"`
	FeeCurrency string  `json:"fee_currency"`
}
