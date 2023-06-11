package Controllers

import (
	"net/http"
	Appservice "varzesh3/AppService"

	//Dto "varzesh3/Dto"

	"github.com/labstack/echo/v4"
)

func GetAllCountries(c echo.Context) error {
	result := Appservice.GetAllCountries()
	//fmt.Println(result)
	c.JSON(http.StatusOK, result)
	return nil
}
