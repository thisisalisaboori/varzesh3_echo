package main

import (
	c "varzesh3/Controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	ec := echo.New()
	ec.GET("countries", c.GetAllCountries)
	ec.GET("history/:id/", c.GetHistory)
	ec.GET("result/:id/", c.GetLeagueResult)
	ec.Start(":8000")
}
