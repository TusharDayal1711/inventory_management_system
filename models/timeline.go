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
