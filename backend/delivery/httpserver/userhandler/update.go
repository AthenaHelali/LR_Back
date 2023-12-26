package userhandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) updateUser(c echo.Context) error {
	var req param.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	cl := claim.GetClaimFromEchoContext(c)
	req.ID = cl.UserID


	res, err := h.userSvc.UpdateUser(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}
