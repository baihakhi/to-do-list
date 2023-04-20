package service

import (
	"context"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/models/payload/request"
	"github.com/baihakhi/to-do-list/internal/repository"
)

type service struct {
	repositories repository.Repository
}

func InitService(repo repository.Repository) Service {
	return &service{
		repositories: repo,
	}
}

type Service interface {
	// User Repository
	CreateUser(context.Context, *models.User) (string, error)
	GetListUser(context.Context, *request.UserRequest) ([]*models.User, error)
	GetOneUser(context.Context, string) (result *models.User, err error)
	UpdateUser(context.Context, *models.User) error
	DeleteUser(context.Context, string) error

	// Activitiy repository
	CreateActivities(context.Context, models.Activity) (int64, error)
	GetListActivities(context.Context, models.GetActivitiesParam) (result []*models.Activity, err error)
	GetOneActivities(context.Context, string, int64) (result *models.Activity, err error)
	UpdateActivities(context.Context, models.UpdateActivityParam) error
	DeleteActivities(context.Context, string, int64) error
}
