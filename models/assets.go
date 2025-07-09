package models

import "time"

type AssetReq struct {
	Brand          string    `json:"brand"`
	Model          string    `json:"model"`
	SerialNo       string    `json:"serial_no"`
	PurchaseDate   time.Time `json:"purchase_date"`
	OwnedBy        string    `json:"owned_by"`
	Type           string    `json:"type"`
	WarrantyStart  time.Time `json:"warranty"`
	WarrantyExpire time.Time `json:"warranty_expire"`
}
