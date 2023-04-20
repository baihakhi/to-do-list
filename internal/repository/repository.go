package repository

import (
	"context"
	"database/sql"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/models/payload/request"
)

type repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

type Repository interface {
	// User Repository
	CreateUser(context.Context, *models.User) (string, error)
	GetListUser(context.Context, *request.PaginationRequest) ([]*models.User, error)
	GetOneUser(context.Context, string) (result *models.User, err error)
	UpdateUser(context.Context, *models.User) error
	DeleteUser(context.Context, string) error

	// Activitiy repository
	CreateActivities(context.Context, models.Activity) (int64, error)
	GetListActivities(context.Context, models.GetActivitiesParam) (result []*models.Activity, err error)
	GetOneActivities(context.Context, int64, string) (result *models.Activity, err error)
	UpdateActivities(context.Context, models.UpdateActivityParam) error
	DeleteActivities(context.Context, string, int64) error
}
