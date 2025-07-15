package models

import "github.com/lib/pq"

type PublicUserReq struct {
	Email string `json:"email"`
}

type EmployeeUpdateReq struct {
	Name      string `json:"name,omitempty"`
	ContactNo string `json:"contact_no,omitempty"`
}

type ManagerRegisterReq struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	ContactNo string `json:"contact_no" validate:"required"`
	Type      string `json:"type" validate:"required,oneof=full_time intern freelancer"`
}

type EmployeeResponseModel struct {
	ID             string         `json:"id" db:"id"`
	Username       string         `json:"username" db:"username"`
	Email          string         `json:"email" db:"email"`
	ContactNo      *string        `json:"contact_no" db:"contact_no"`
	EmployeeType   string         `json:"type" db:"employee_type"`
	AssignedAssets pq.StringArray `json:"assigned_assets" db:"assigned_assets"`
}

type UpdateUserRoleReq struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	Role   string `json:"role" validate:"required,oneof=admin inventory_manager employee_manager user"`
}

type UpdateUserReq struct {
	UserID    string  `json:"user_id" validate:"required,uuid"`
	Username  *string `json:"username,omitempty"`
	ContactNo *string `json:"contact_no,omitempty"`
	UpdatedBy string  `json:"-"`
}
