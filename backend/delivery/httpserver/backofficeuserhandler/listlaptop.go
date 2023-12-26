package backofficeuserhandler

import (
	"game-app/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) listLaptops(c echo.Context) error {
	list, err := h.backofficeUserSvc.ListLaptops()
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": list,
	})
}