package userhandler

import (
	"game-app/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUerRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)

	userGroup.POST("/login", h.userLogin)

	userGroup.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))
	
	userGroup.POST("/:user_id/favorites",h.addFavorite,middleware.Auth(h.authSvc, h.authConfig))
	
	userGroup.GET("/:user_id/favorites",h.getLaptops,middleware.Auth(h.authSvc, h.authConfig))

	userGroup.POST("/:user_id/search",h.search,middleware.Auth(h.authSvc, h.authConfig))


}
