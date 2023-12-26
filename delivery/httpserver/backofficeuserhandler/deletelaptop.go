package backofficeuserhandler

import (
	"game-app/param"
	"game-app/pkg/httpmsg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
// @Summary Delete laptop
// @Description Delete a laptop by ID
// @Tags backoffice
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Security BearerAuth
// @Param laptop_id path integer true "ID of the laptop to be deleted"
// @Success 200 "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Laptop not found"
// @Router /backoffice/laptops/{laptop_id} [delete]
func (h Handler) deletelaptop(c echo.Context) error {
	var dReq param.DeleteLaptopRequest
	laptop_Id,_ := strconv.Atoi(c.Param("laptop_id"))
	dReq.LaptopID =uint64(laptop_Id)
	err := h.backofficeUserSvc.DeleteLaptop(dReq)
	if err != nil {
		msg, code := httpmsg.HTTPCodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK,nil)
}