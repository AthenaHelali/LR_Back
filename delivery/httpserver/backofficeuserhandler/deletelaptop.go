package backofficeuserhandler

import (
	"game-app/param"
	"game-app/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)
func (h Handler) deletelaptop(c echo.Context) error {
	var dReq param.DeleteLaptopRequest

	if err := c.Bind(&dReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}
	err := h.backofficeUserSvc.DeleteLaptop(dReq)
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK,nil)
}