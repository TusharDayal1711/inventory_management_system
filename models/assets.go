package models

import "time"

type AssetReq struct {
	ID             string    `json:"id"`
	Brand          string    `json:"brand"`
	Model          string    `json:"model"`
	SerialNo       string    `json:"serial_no"`
	PurchaseDate   time.Time `json:"purchase_date"`
	OwnedBy        string    `json:"owned_by"`
	Type           string    `json:"type"`
	WarrantyStart  time.Time `json:"warranty"`
	WarrantyExpire time.Time `json:"warranty_expire"`
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

type AssetConfigReq struct {
	ID             string `json:"id"`
	Processor      string `json:"processor,omitempty"`
	Ram            string `json:"ram,omitempty"`
	Os             string `json:"os,omitempty"`
	DPI            string `json:"dpi,omitempty"`
	Display        string `json:"display,omitempty"`
	Resolution     string `json:"resolution,omitempty"`
	Port           string `json:"port,omitempty"`
	HDDType        string `json:"hdd_type,omitempty"`
	Storage        string `json:"storage,omitempty"`
	Version        string `json:"version,omitempty"`
	IMEI1          string `json:"imei1,omitempty"`
	IMEI2          string `json:"imei2,omitempty"`
	Number         int    `json:"number,omitempty"`
	AccessoryType  string `json:"accessory_type,omitempty"`
	AdditionalInfo string `json:"additional_info,omitempty"`
}

var TypeTableName = map[string]string{
	"laptop":     "laptop_config_Req",
	"mouse":      "mouse_config_Req",
	"monitor":    "monitor_config_Req",
	"hard_disk":  "hard_disk_config_Req",
	"pen_driver": "pen_driver_config_Req",
	"mobile":     "mobile_config_Req",
	"sim":        "sim_config_Req",
	"accessory":  "accessory_config_Req",
}
