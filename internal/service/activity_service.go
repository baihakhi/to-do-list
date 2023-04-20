package service

import (
	"context"

	"github.com/baihakhi/to-do-list/internal/models"
)

// CreateActivities insert activity data into database
// return activityID if success
func (s *service) CreateActivities(ctx context.Context, data models.Activity) (int64, error) {
	return s.repositories.CreateActivities(ctx, data)
}

// GetListActivities get array of activity data(s) from database
func (s *service) GetListActivities(ctx context.Context, param models.GetActivitiesParam) ([]*models.Activity, error) {
	return s.repositories.GetListActivities(ctx, param)
}

// GetOneActivities get activity data from database
func (s *service) GetOneActivities(ctx context.Context, username string, activityID int64) (*models.Activity, error) {
	return s.repositories.GetOneActivities(ctx, activityID, username)
}

// UpdateActivities update data activity in database
func (s *service) UpdateActivities(ctx context.Context, data models.UpdateActivityParam) error {
	return s.repositories.UpdateActivities(ctx, data)
}

// DeleteActivities delete activity data in database
func (s *service) DeleteActivities(ctx context.Context, username string, activityID int64) error {
	return s.repositories.DeleteActivities(ctx, username, activityID)
}
