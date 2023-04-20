package router

import (
	"net/http"

	"github.com/baihakhi/to-do-list/internal/handler"
	"github.com/baihakhi/to-do-list/internal/middleware"
	"github.com/baihakhi/to-do-list/internal/models/enum"
	"github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Internal API TODO App!")
	})

	v1 := server.Group("/api/v1")
	{
		user := v1.Group("/user")
		{

			user.POST("", handler.CreateUser,
				middleware.SetMiddlewareAuthentication([]string{enum.RoleSuperAdmin}))
			user.GET("", handler.GetListUser, middleware.SetMiddlewareAuthentication(enum.AllAccess))
			user.GET("/:username", handler.GetOneUser, middleware.SetMiddlewareAuthentication(enum.AllAccess))
			user.PUT("", handler.EditUser, middleware.SetMiddlewareAuthentication([]string{enum.RoleSuperAdmin}))
			user.DELETE("", handler.RemoveUser, middleware.SetMiddlewareAuthentication([]string{enum.RoleSuperAdmin}))

		}
	}
}
