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

	if !datasource.IsValidSymbol(s) {
		c.JSON(http.StatusNotFound, "Invalid Symbol.")
		return
	}
	if !datasource.IsValidSymbolSupported(s) {
		c.JSON(http.StatusNotFound, "Symbol not supported.")
		return
	}

	//Get currency from in memory store
	d := domain.Currency{
		ID: "ETH",
	}
	c.JSON(http.StatusOK, d)
}

//GetAll returns all symbol data.
func GetAll(c *vodka.Context) {
	s := c.PathParams["symbol"]
	if s == "" {
		c.JSON(http.StatusNotFound, s)
	}

	if datasource.IsValidSymbol(s) {
		c.JSON(http.StatusNotFound, s)
	}

	d := domain.Currency{
		ID: "ETH",
	}
	c.JSON(http.StatusOK, d)
}
