package controllers

import (
	"net/http"

	"github.com/smudra1990/goCodeTest/domain"
	"github.com/smudra1990/goCodeTest/vodka"
)

//GetAll is health check endpoint for microservice.
func GetAll(c *vodka.Context) {
	s := c.PathParams["symbol"]
	if c == nil {
		c.JSON(http.StatusNotFound, s)
	}

	d := domain.Domain{
		ID: "ETH",
	}
	c.JSON(http.StatusOK, d)
}
