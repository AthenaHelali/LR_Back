package backofficeuserhandler

import (
	"game-app/param"
	"game-app/pkg/httpmsg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
func (h Handler) deletelaptop(c echo.Context) error {
	var dReq param.DeleteLaptopRequest
	laptop_Id,_ := strconv.Atoi(c.Param("laptop_id"))
	dReq.LaptopID =uint64(laptop_Id)
	err := h.backofficeUserSvc.DeleteLaptop(dReq)
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK,nil)
}