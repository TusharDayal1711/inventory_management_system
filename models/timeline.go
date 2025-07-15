package models

import (
	"time"
)

type UserTimelineRes struct {
	AssetID      string     `json:"asset_id" db:"asset_id"`
	Brand        string     `json:"brand" db:"brand"`
	Model        string     `json:"model" db:"model"`
	SerialNo     string     `json:"serial_no" db:"serial_no"`
	AssignedAt   time.Time  `json:"assigned_at" db:"assigned_at"`
	ReturnedAt   *time.Time `json:"returned_at,omitempty" db:"returned_at"`
	ReturnReason *string    `json:"return_reason,omitempty" db:"return_reason"`
}

type AssetTimelineRes struct {
	EmployeeID   string     `json:"employee_id" db:"employee_id"`
	EmployeeName string     `json:"employee_name" db:"employee_name"`
	Email        string     `json:"email" db:"email"`
	ContactNo    *string    `json:"contact_no,omitempty" db:"contact_no"`
	AssignedAt   time.Time  `json:"assigned_at" db:"assigned_at"`
	ReturnedAt   *time.Time `json:"returned_at,omitempty" db:"returned_at"`
	ReturnReason *string    `json:"return_reason,omitempty" db:"return_reason"`
}

type DeleteUserReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
}
