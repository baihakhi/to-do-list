package request

import (
	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/labstack/echo/v4"
)

type ActivityRequest struct {
	PaginationRequest
	UserAccess
	models.Activity
}

func (g *ActivityRequest) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(g); err != nil {
		return err
	}

	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, g)

	g.ValidatePagination()

	return nil
}
