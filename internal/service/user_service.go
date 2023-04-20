package service

import (
	"context"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/models/payload/request"
)

// CreateUser insert user data into database
// return userID if success
func (s *service) CreateUser(ctx context.Context, data *models.User) (string, error) {
	return s.repositories.CreateUser(ctx, data)
}

// GetListUser get array of user data(s) from database
func (s *service) GetListUser(ctx context.Context, param *request.UserRequest) ([]*models.User, error) {
	return s.repositories.GetListUser(ctx, &param.PaginationRequest)
}

// GetOneUser get user data from database
func (s *service) GetOneUser(ctx context.Context, param string) (*models.User, error) {
	return s.repositories.GetOneUser(ctx, param)
}

// UpdateUser update data user in database
func (s *service) UpdateUser(ctx context.Context, data *models.User) error {
	return s.repositories.UpdateUser(ctx, data)
}

// DeleteUser delete user data in database
func (s *service) DeleteUser(ctx context.Context, param string) error {
	return s.repositories.DeleteUser(ctx, param)
}
