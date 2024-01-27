package userhandler

import (
	"game-app/param"
	"game-app/pkg/claim"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary Add a laptop to favorites
// @Description Add a laptop to the user's favorites list
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param request body param.AddLaptopRequest true "Laptop information to add to favorites"
// @Success 200 {object} param.AddLaptopResponse "Successfully added laptop to favorites"
// @Failure 400 {object} error "Bad Request"
// @Router /users/laptops/favorites [post]
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

// @Summary Remove laptop from favorites
// @Description Remove a laptop from the user's favorites list
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param laptop_id path integer true "ID of the laptop to be removed from favorites"
// @Param request body param.RemoveFavoriteLaptopRequest true "Request parameters for removing a laptop from favorites"
// @Success 200 {object} param.RemoveFavoriteLaptopResponse "Successfully removed laptop from favorites"
// @Failure 400 {object} error "Bad Request"
// @Router /users/laptops/favorites/{laptop_id} [delete]
func (h Handler) removeFavorite(c echo.Context) error {
	var req param.RemoveFavoriteLaptopRequest
	cl := claim.GetClaimFromEchoContext(c)
	laptop_Id, _ := strconv.Atoi(c.Param("laptop_id"))
	req.LaptopID = laptop_Id

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

// @Summary Get user's favorite laptops
// @Description Retrieve the list of laptops marked as favorites by the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param request query param.LaptopsRequest false "Query parameters for retrieving favorite laptops"
// @Success 200 {object} error "List of favorite laptops"
// @Failure 400 {object} error "Bad Request"
// @Router /users/laptops/favorites [get]
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

// @Summary Get laptop details by ID
// @Description Retrieve information about a laptop using its ID.
// @Tags laptops
// @Accept json
// @Produce json
// @Param laptop_id path int true "Laptop ID" Format(int64)
// @Success 200 {object} error "Successful response"
// @Failure 400 {object} error "Bad Request"
// @Router /userGroup/laptop/{laptop_id} [get]
func (h Handler) getLaptop(c echo.Context) error {
	var req param.LaptopRequest
	laptop_Id, _ := strconv.Atoi(c.Param("laptop_id"))
	req.LaptopID = laptop_Id

	response, err := h.userSvc.GetLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}

// @Summary Search for laptops
// @Description Perform a search for laptops based on the provided criteria
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param request body param.SearchRequest true "Search criteria for laptops"
// @Success 200 {object} param.SearchResponse "Search results for laptops"
// @Failure 400 {object} error "Bad Request"
// @Router /users/laptops/search [post]
func (h Handler) search(c echo.Context) error {
	var req param.SearchRequest

	cl := claim.GetClaimFromEchoContext(c)
	req.UserID = int(cl.UserID)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	response, err := h.userSvc.Search(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
// @Summary Update a laptop for a seller
// @Description Allows a seller to update details of a laptop in their inventory.
// @Tags sellers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param request body param.UpdateLaptopRequest true "Updated laptop details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} error "Bad Request"
// @Router /sellerGroup/laptops/ [put]
func (h Handler) updateLaptop(c echo.Context) error {
	var req param.UpdateLaptopRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	res, err := h.userSvc.UpdateLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
// @Summary Remove a laptop for a seller
// @Description Allows a seller to remove a laptop from their inventory.
// @Tags sellers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param laptop_id path int true "Laptop ID" Format(int64)
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} error "Bad Request"
// @Router /sellerGroup/laptops/{laptop_id} [delete]
func (h Handler) removeSellerLaptop(c echo.Context) error {
	var req param.RemoveSellerLaptopRequest
	cl := claim.GetClaimFromEchoContext(c)
	laptop_Id, _ := strconv.Atoi(c.Param("laptop_id"))
	req.LaptopID = laptop_Id

	req.UserID = int(cl.UserID)

	resp, err := h.userSvc.RemoveSellerLaptop(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
