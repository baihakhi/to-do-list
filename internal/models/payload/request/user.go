package request

import (
	"net/mail"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/models/enum"
	"github.com/baihakhi/to-do-list/internal/models/payload/response"
	"github.com/labstack/echo/v4"
)

type (
	UserRequest struct {
		PaginationRequest
		UserAccess
		models.User
	}
)

// GetDataFromHTTPRequest bind from HTTP Request to UserRequest struct
// using echo binder
func (u *UserRequest) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, u)

	u.ValidatePagination()

	return nil
}

// Validate validate UserRequest struct
func (u *UserRequest) Validate() error {
	if !enum.IsValidUserGroup[u.UserGroup] {
		return response.ErrorBuilder(enum.UGroup, response.UNKN)
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return response.ErrorBuilder(enum.UMail, response.INVL)
	}
	return nil
}
