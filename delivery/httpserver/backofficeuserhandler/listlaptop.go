package backofficeuserhandler

import (
	"game-app/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)
// @Summary List laptops
// @Description Retrieve a list of laptops
// @Tags backoffice
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "List of laptops"
// @Failure 400 {object} error "Bad Request"
// @Router /backoffice/laptops [get]
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