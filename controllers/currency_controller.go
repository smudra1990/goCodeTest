package controllers

import (
	"net/http"

	datasource "github.com/smudra1990/goCodeTest/datasource/currencydb"
	"github.com/smudra1990/goCodeTest/domain"
	"github.com/smudra1990/goCodeTest/vodka"
)

//Get specific symbol data.
func Get(c *vodka.Context) {
	s := c.PathParams["symbol"]
	sym := datasource.GetSymbol(s)

	if !datasource.IsValidSymbolSupported(s) {
		c.JSON(http.StatusNotFound, "Symbol not supported.")
		return
	}

	if sym == nil {
		c.JSON(http.StatusNotFound, "Invalid Symbol.")
		return
	}

	d := domain.Currency{
		ID:          sym.BaseCurrency,
		FullName:    sym.FullName,
		Ask:         sym.Params.Ask,
		Bid:         sym.Params.Bid,
		Last:        sym.Params.Last,
		Open:        sym.Params.Open,
		High:        sym.Params.High,
		Low:         sym.Params.Low,
		FeeCurrency: sym.FeeCurrency,
	}
	c.JSON(http.StatusOK, d)
}

//GetAll returns all symbol data.
func GetAll(c *vodka.Context) {
	result := []domain.Currency{}
	allSym := datasource.GetAllSymbol()
	for _, sym := range *allSym {
		d := domain.Currency{
			ID:          sym.BaseCurrency,
			FullName:    sym.FullName,
			Ask:         sym.Params.Ask,
			Bid:         sym.Params.Bid,
			Last:        sym.Params.Last,
			Open:        sym.Params.Open,
			High:        sym.Params.High,
			Low:         sym.Params.Low,
			FeeCurrency: sym.FeeCurrency,
		}
		result = append(result, d)
	}
	c.JSON(http.StatusOK, result)
}
