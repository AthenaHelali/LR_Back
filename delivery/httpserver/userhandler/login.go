package userhandler

import (
	"game-app/param"
	"game-app/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary User login
// @Description Login with the provided credentials
// @Tags users
// @Accept json
// @Produce json
// @Param request body param.LoginRequest true "User login information"
// @Success 200 {object} param.LoginResponse "Successfully logged in"
// @Failure 400 {object} error "Bad Request"
// @Failure 401 {object} error "Unauthorized"
// @Router /users/login [post]
func (h Handler) userLogin(c echo.Context) error {
	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	if err := h.userValidator.ValidateLoginRequest(req); err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	response, err := h.userSvc.Login(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
