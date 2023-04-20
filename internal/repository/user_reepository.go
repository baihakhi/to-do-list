package repository

import (
	"context"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/models/payload/request"
	"github.com/baihakhi/to-do-list/internal/repository/queries"
)

// CreateUser do query insert user data into database
// return user code if success
func (a *repository) CreateUser(ctx context.Context, data *models.User) (string, error) {
	var userID string
	result := a.db.QueryRowContext(ctx, queries.CreateUsers,
		data.Code,
		data.Name,
		data.UserGroup,
		data.Email,
		data.Password,
	)

	if err := result.Err(); err != nil {
		return "", err
	}
	result.Scan(&userID)

	return userID, nil
}

// GetListUser do query get array of user data(s) from database
func (a *repository) GetListUser(ctx context.Context, params *request.PaginationRequest) (result []*models.User, err error) {
	rows, err := a.db.QueryContext(ctx, queries.GetListUsers, params.Limit, (params.Page-1)*params.Limit, params.Key)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var data models.User
		if rows.Scan(
			&data.ID,
			&data.Code,
			&data.Name,
			&data.UserGroup,
			&data.Email,
			&data.Status,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, &data)
	}

	return result, nil
}

// GetOneUser do query get user data from database
func (a *repository) GetOneUser(ctx context.Context, param string) (*models.User, error) {
	rows := a.db.QueryRowContext(ctx, queries.GetOneUsersByCode, param)

	var data models.User
	err := rows.Scan(
		&data.ID,
		&data.Code,
		&data.Name,
		&data.UserGroup,
		&data.Email,
		&data.Status,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return &data, err
}

// UpdateUser do query update data user in database
func (a *repository) UpdateUser(ctx context.Context, params *models.User) error {
	_, err := a.db.ExecContext(ctx, queries.UpdateActivities,
		params.Code,
		params.Name,
		params.UserGroup,
		params.Email,
		params.Password,
	)

	return err
}

// DeleteUser do query delete user data in database
func (a *repository) DeleteUser(ctx context.Context, username string) error {
	_, err := a.db.ExecContext(ctx, queries.DeleteActivities, username)

	return err
}
