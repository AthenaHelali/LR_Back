package userhandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"game-app/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)
// @Summary Get user profile
// @Description Retrieve the profile information of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} param.ProfileResponse "User profile information"
// @Failure 401 {object} error "Unauthorized"
// @Router /users/profile [get]
func (h Handler) userProfile(c echo.Context) error {

	cl := claim.GetClaimFromEchoContext(c)

	resp, err := h.userSvc.Profile(param.ProfileRequest{UserID: cl.UserID})
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)

}
