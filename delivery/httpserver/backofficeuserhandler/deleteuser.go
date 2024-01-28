package backofficeuserhandler

import (
	"game-app/param"
	"game-app/pkg/httpmsg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) deleteuser(c echo.Context) error {
	var dReq param.DeleteUserRequest
	user_Id, _ := strconv.Atoi(c.Param("user_id"))
	dReq.UserID = user_Id
	resp, err := h.backofficeUserSvc.DeleteUser(dReq)
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
