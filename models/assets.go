package models

import (
	"encoding/json"
	"time"
)

type AssetReq struct {
	Brand          string    `json:"brand" validate:"required"`
	Model          string    `json:"model" validate:"required"`
	SerialNo       string    `json:"serial_no" validate:"required"`
	PurchaseDate   time.Time `json:"purchase_date" validate:"required"`
	OwnedBy        string    `json:"owned_by" validate:"required"`
	Type           string    `json:"type" validate:"required"`
	WarrantyStart  time.Time `json:"warranty" validate:"required"`
	WarrantyExpire time.Time `json:"warranty_expire" validate:"required,gtfield=WarrantyStart"`
}

// Assets request model
type Laptop_config_req struct {
	SerialNo  string `json:"serial_no"`
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Os        string `json:"os"`
}
type Mouse_config_req struct {
	SerialNo string `json:"serial_no"`
	DPI      string `json:"dpi"`
}

type Monitor_config_req struct {
	SerialNo   string `json:"serial_no"`
	Display    string `json:"display"`
	Resolution string `json:"resolution"`
	Port       string `json:"port"`
}

type Hard_disk_config_req struct {
	SerialNo string `json:"serial_no"`
	Type     string `json:"type"`
	Storage  string `json:"storage"`
}

type Pen_drive_config_req struct {
	SerialNo string `json:"serial_no"`
	Version  string `json:"version"`
	Storage  string `json:"storage"`
}

type Mobile_config_req struct {
	SerialNo  string `json:"serial_no"`
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Os        string `json:"os"`
	IMEI1     string `json:"imei"`
	IMEI2     string `json:"ime2"`
}

type Sim_config_req struct {
	SerialNo string `json:"serial_no"`
	Number   int    `json:"number"`
}

type Accessories_config_req struct {
	SerialNo       string `json:"serial_no"`
	Type           string `json:"type"`
	AdditionalInfo string `json:"additional_info"`
}

type AddAssetWithConfigReq struct {
	AssetReq
	Config json.RawMessage `json:"config"`
}
