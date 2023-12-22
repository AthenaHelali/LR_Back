package userhandler

import (
	"game-app/param"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func(h Handler)updateUser(c echo.Context) error{
	var req param.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}
	userId,_ := strconv.Atoi(c.Param("user_id"))
	req.ID = uint(userId)
	res, err := h.userSvc.UpdateUser(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
	


}
