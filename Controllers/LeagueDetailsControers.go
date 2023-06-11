package Controllers

import (

	//Dto "varzesh3/Dto"

	"net/http"
	"strconv"
	Appservice "varzesh3/AppService"

	"github.com/labstack/echo/v4"
)

func GetLeagueResult(c echo.Context) error {
	//fmt.Println("Result")
	id, _ := strconv.Atoi(c.Param("id"))
	result := Appservice.GetLeagueResult(id)
	c.JSON(http.StatusOK, result)
	return nil
}
