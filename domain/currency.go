package domain

// Currency struct
type Currency struct {
	ID          string `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Ask         string `json:"ask,omitempty"`
	Bid         string `json:"bid,omitempty"`
	Last        string `json:"last,omitempty"`
	Open        string `json:"open,omitempty"`
	Low         string `json:"low,omitempty"`
	High        string `json:"high,omitempty"`
	FeeCurrency string `json:"fee_currency,omitempty"`
}
