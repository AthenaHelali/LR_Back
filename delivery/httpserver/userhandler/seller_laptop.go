package userhandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get laptops for a seller
// @Description Retrieve the list of laptops associated with the authenticated seller.
// @Tags sellers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} param.LaptopsResponse "Successful response"
// @Failure 400 {object} error "Bad Request"
// @Router /sellerGroup/laptops/ [get]
func (h Handler) getSellerLaptops(c echo.Context) error {
	var req param.LaptopsRequest

	cl := claim.GetClaimFromEchoContext(c)

	req.UserID = int(cl.UserID)

	response, err := h.userSvc.GetSellerLaptops(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
