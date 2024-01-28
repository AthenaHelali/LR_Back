package userhandler

import (
	"fmt"
	"game-app/param"
	"game-app/pkg/claim"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Update user information
// @Description Update user information with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param request body param.UpdateUserRequest true "User update information"
// @Success 200 {object} param.UpdateUserResponse "Successfully updated user"
// @Failure 400 {object} error "Bad Request"
// @Failure 401 {object} error "Unauthorized"
// @Failure 404 {object} error "User not found"
// @Router /users/ [put]
func (h Handler) updateUser(c echo.Context) error {
	var req param.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	cl := claim.GetClaimFromEchoContext(c)
	req.ID = cl.UserID

	fmt.Println(req)
	res, err := h.userSvc.UpdateUser(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}
