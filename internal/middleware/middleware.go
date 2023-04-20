package middleware

import (
	"context"
	"net/http"

	"github.com/baihakhi/to-do-list/internal/models/enum"
	"github.com/baihakhi/to-do-list/internal/models/payload/response"
	"github.com/baihakhi/to-do-list/internal/utils"
	"github.com/labstack/echo/v4"
)

// SetMiddlewareAuthentication verify user token
func SetMiddlewareAuthentication(access []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := TokenValid(c.Request())
			if err != nil {
				c.JSON(http.StatusUnauthorized, response.MapResponse{
					Message: response.InvalidToken,
				})
				return err
			}

			isAllowed := utils.IsContain(user.UserGroup, access)
			if !isAllowed {
				c.JSON(http.StatusForbidden, response.MapResponse{
					Message: response.AccessDenied,
				})
				return err
			}
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), enum.ACC, user)))
			return next(c)
		}
	}
}
