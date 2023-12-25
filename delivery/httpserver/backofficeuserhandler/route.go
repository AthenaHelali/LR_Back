package backofficeuserhandler

import (
	"game-app/delivery/httpserver/middleware"
	"game-app/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetBackOfficeUerRoutes(e *echo.Echo) {
	h.backofficeUserSvc.RegisterAdmin()

	userGroup := e.Group("/backoffice")

	userGroup.GET("/users/", h.listUsers, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))

	// userGroup.DELETE("/users/", h.listUsers, middleware.Auth(h.authSvc, h.authConfig),
	// 	middleware.AccessCheck(h.authorizationSvc, entity.UserDeletePermission))

	userGroup.GET("/laptops/", h.listLaptops, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.LaptopListPermission))

	userGroup.DELETE("/laptops/", h.deletelaptop, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))

}
