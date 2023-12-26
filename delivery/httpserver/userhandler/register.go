package userhandler

import (
	"game-app/param"
	"game-app/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param request body param.RegisterRequest true "User registration information"
// @Success 201 {object} param.RegisterResponse "Successfully registered"
// @Failure 400 {object} error "Bad Request"
// @Router /users/register [post]
func (h Handler) userRegister(c echo.Context) error {
	var uReq param.RegisterRequest

	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}
	if err := h.userValidator.ValidateRegisterRequest(uReq); err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	response, err := h.userSvc.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}
