package models

import "github.com/baihakhi/to-do-list/internal/models/enum"

type (
	User struct {
		Identity
		UserGroup string            `form:"user_group" json:"user_group"`
		Email     string            `form:"email" json:"email"`
		Password  string            `form:"password" json:"password,omitempty"`
		Token     string            `json:"tokenpassword,omitempty"`
		Status    enum.BinaryStatus `form:"status" json:"status"`
		Timestamp
	}
)
