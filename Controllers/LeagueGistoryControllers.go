package Controllers

import (

	//Dto "varzesh3/Dto"

	"fmt"
	"net/http"
	"strconv"
	Appservice "varzesh3/AppService"

	"github.com/labstack/echo/v4"
)

func GetHistory(c echo.Context) error {
	fmt.Println("History")
	id, _ := strconv.Atoi(c.Param("id"))
	result := Appservice.GetHistory(id)
	c.JSON(http.StatusOK, result)
	return nil
}
