package userhandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"net/http"

	"github.com/labstack/echo/v4"
)


func (h Handler)getSellerLaptops(c echo.Context) error{
	var req param.LaptopsRequest

	cl := claim.GetClaimFromEchoContext(c)

	req.UserID = int(cl.UserID)

	response, err := h.userSvc.GetSellerLaptops(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}