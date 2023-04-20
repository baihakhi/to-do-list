package repository

import (
	"context"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/repository/queries"
)

// CreateActivities do query insert activities data into database
// return activitiesID if success
func (r *repository) CreateActivities(ctx context.Context, data models.Activity) (int64, error) {
	result, err := r.db.ExecContext(ctx, queries.CreateActivities,
		data.Code,
		data.Name,
		data.Description,
		data.Assignee,
		data.Owner)

	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// GetListActivities do query get array of activities data(s) from database
func (r *repository) GetListActivities(ctx context.Context, params models.GetActivitiesParam) (result []*models.Activity, err error) {
	rows, err := r.db.QueryContext(ctx, queries.GetListActivities, params.Owner, params.Limit, (params.Offset-1)*params.Limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var data models.Activity
		if rows.Scan(
			&data.ID,
			&data.Code,
			&data.Name,
			&data.Description,
			&data.Owner,
			&data.Assignee,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, &data)
	}

	return result, nil
}

// GetOneActivities do query get activities data from database
func (r *repository) GetOneActivities(ctx context.Context, actID int64, activitiesname string) (result *models.Activity, err error) {
	rows := r.db.QueryRowContext(ctx, queries.GetOneActivitiesByID, actID, activitiesname)

	var data models.Activity
	err = rows.Scan(
		&data.ID,
		&data.Code,
		&data.Name,
		&data.Description,
		&data.Owner,
		&data.Assignee,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return result, err
}

// UpdateActivities do query update data activities in database
func (r *repository) UpdateActivities(ctx context.Context, params models.UpdateActivityParam) error {
	_, err := r.db.ExecContext(ctx, queries.UpdateActivities,
		params.Code,
		params.Name,
		params.Description,
		params.Owner,
		params.Assignee,
		params.ID,
		params.Username,
	)

	return err
}

// DeleteActivities do query delete activities data in database
func (r *repository) DeleteActivities(ctx context.Context, activitiesname string, activityID int64) error {
	_, err := r.db.ExecContext(ctx, queries.DeleteActivities, activityID, activitiesname)

	return err
}
