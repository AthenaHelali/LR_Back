package userhandler

import (
	"game-app/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetUerRoutes(e *echo.Echo) {	
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)

	userGroup.POST("/login", h.userLogin)

	userGroup.PUT("/", h.updateUser, middleware.Auth(h.authSvc, h.authConfig))

	userGroup.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))

	userGroup.POST("/laptops/favorites", h.addFavorite, middleware.Auth(h.authSvc, h.authConfig))

	userGroup.GET("/laptops/favorites", h.getLaptops, middleware.Auth(h.authSvc, h.authConfig))

	userGroup.DELETE("/laptops/favorites/:laptop_id", h.removeFavorite, middleware.Auth(h.authSvc, h.authConfig))

	userGroup.POST("/laptops/search", h.search, middleware.Auth(h.authSvc, h.authConfig))

	userGroup.GET("/laptop/:laptop_id", h.getLaptop)

	sellerGroup := e.Group("/sellers")

	sellerGroup.POST("/laptops/", h.AddSellerLaptop, middleware.Auth(h.authSvc, h.authConfig))

	sellerGroup.GET("/laptops/", h.getSellerLaptops, middleware.Auth(h.authSvc, h.authConfig))

	sellerGroup.POST("/upload/", h.uploadImage)

	sellerGroup.PUT("/laptops/", h.updateLaptop, middleware.Auth(h.authSvc, h.authConfig))

	sellerGroup.DELETE("/laptops/:laptop_id", h.removeSellerLaptop, middleware.Auth(h.authSvc, h.authConfig))

}
