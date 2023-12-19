package userhandler

import (
	"game-app/param"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) addFavorite(c echo.Context) error {
	var req param.AddLaptopRequest

	userId := c.Param("user_id")

	req.UserID,_ = strconv.Atoi(userId)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}


	response, err := h.userSvc.AddFavoriteLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}

func (h Handler) getLaptops(c echo.Context) error {
	var req param.LaptopsRequest

	userId := c.Param("user_id")

	req.UserID,_ = strconv.Atoi(userId)


	response, err := h.userSvc.GetLaptops(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
func (h Handler) search(c echo.Context) error {
	var req param.SearchRequest

	userId := c.Param("user_id")
	req.UserID,_ = strconv.Atoi(userId)

	response, err := h.userSvc.Search(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)


}
