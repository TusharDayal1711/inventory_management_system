package models

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

type EmployeeAssetRes struct {
	ID            string   `json:"id"`
	Username      string   `json:"name"`
	Email         string   `json:"email"`
	ContactNo     string   `json:"contact_no"`
	Type          string   `json:"type"`
	AssetAssigned []string `json:"asset_assigned"`
}
