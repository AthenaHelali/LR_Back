package userhandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) addFavorite(c echo.Context) error {
	var req param.AddLaptopRequest
	cl := claim.GetClaimFromEchoContext(c)

	req.UserID = int(cl.UserID)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	response, err := h.userSvc.AddFavoriteLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}

func (h Handler) removeFavorite(c echo.Context) error {
	var req param.RemoveFavoriteLaptopRequest
	cl := claim.GetClaimFromEchoContext(c)
	laptop_Id,_ := strconv.Atoi(c.Param("laptop_id"))
	req.LaptopID =laptop_Id

	req.UserID = int(cl.UserID)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	resp, err := h.userSvc.RemoveFavoriteLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}

func (h Handler) getLaptops(c echo.Context) error {
	var req param.LaptopsRequest

	cl := claim.GetClaimFromEchoContext(c)

	req.UserID = int(cl.UserID)

	response, err := h.userSvc.GetLaptops(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}

func (h Handler) getLaptop(c echo.Context) error {
	var req param.LaptopRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	response, err := h.userSvc.GetLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
func (h Handler) search(c echo.Context) error {
	var req param.SearchRequest

	cl := claim.GetClaimFromEchoContext(c)
	req.UserID = int(cl.UserID)

	response, err := h.userSvc.Search(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
