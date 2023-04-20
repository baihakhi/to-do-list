package handler

import (
	"net/http"

	"github.com/baihakhi/to-do-list/internal/models/payload/request"
	"github.com/baihakhi/to-do-list/internal/models/payload/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	data := new(request.UserRequest)
	ctx := c.Request().Context()

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	result, err := h.service.CreateUser(ctx, &data.User)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.SUCCESS,
		Data: map[string]string{
			"user_id": result,
		},
	})
}

func (h *Handler) GetListUser(c echo.Context) error {
	data := new(request.UserRequest)
	ctx := c.Request().Context()

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	result, err := h.service.GetListUser(ctx, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	pagination := &response.Pagination{
		Page:      data.Page,
		Limit:     data.Limit,
		TotalData: len(result),
	}
	pagination.CountTotalPage()

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message:  response.SUCCESS,
		Data:     result,
		Metadata: pagination,
	})
}

func (h *Handler) GetOneUser(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.Param("username")
	result, err := h.service.GetOneUser(ctx, code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.SUCCESS,
		Data:    result,
	})
}

func (h *Handler) EditUser(c echo.Context) error {
	data := new(request.UserRequest)
	ctx := c.Request().Context()

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	err := h.service.UpdateUser(ctx, &data.User)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.SUCCESS,
		Data:    data.ID,
	})
}

func (h *Handler) RemoveUser(c echo.Context) error {
	data := new(request.UserRequest)
	ctx := c.Request().Context()

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	err := h.service.DeleteUser(ctx, data.Code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.SUCCESS,
		Data:    data.ID,
	})
}
