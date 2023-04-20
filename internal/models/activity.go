package models

import (
	"database/sql"
	"errors"
)

type (
	Activity struct {
		Identity
		Description sql.NullString `json:"description"`
		Owner       string         `json:"owner"`
		Assignee    sql.NullString `json:"assignee"`
		Timestamp
	}

	GetActivitiesParam struct {
		Owner         string
		Limit, Offset int8
	}

	UpdateActivityParam struct {
		Activity
		Username string
	}
)

// Validate do validation for activity
func (a *Activity) Validate() error {
	if a.Name == "" {
		return errors.New("")
	}
	return nil
}
