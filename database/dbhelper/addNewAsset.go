package dbhelper

import (
	"fmt"
	"inventory_management_system/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func AddAsset(tx *sqlx.Tx, assetReq models.AddAssetWithConfigReq, addedBy uuid.UUID) (uuid.UUID, error) {
	var assetID uuid.UUID
	err := tx.Get(&assetID, `
		INSERT INTO assets (
			brand, model, serial_no, purchase_date, 
			owned_by, type, warranty_start, warranty_expire, 
			added_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`,
		assetReq.Brand, assetReq.Model, assetReq.SerialNo, assetReq.PurchaseDate,
		assetReq.OwnedBy, assetReq.Type, assetReq.WarrantyStart, assetReq.WarrantyExpire,
		addedBy)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert asset: %w", err)
	}
	return assetID, nil
}

func AddLaptopConfig(tx *sqlx.Tx, cfg models.Laptop_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO laptop_config (asset_id, processor, ram, os)
		VALUES ($1, $2, $3, $4)`,
		assetID, cfg.Processor, cfg.Ram, cfg.Os)
	if err != nil {
		return fmt.Errorf("failed to insert laptop config: %w", err)
	}
	return nil
}

func AddMouseConfig(tx *sqlx.Tx, cfg models.Mouse_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO mouse_config (asset_id, dpi)
		VALUES ($1, $2)`,
		assetID, cfg.DPI)
	if err != nil {
		return fmt.Errorf("failed to insert mouse config: %w", err)
	}
	return nil
}

func AddMonitorConfig(tx *sqlx.Tx, cfg models.Monitor_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO monitor_config (asset_id, display, resolution, port)
		VALUES ($1, $2, $3, $4)`,
		assetID, cfg.Display, cfg.Resolution, cfg.Port)
	if err != nil {
		return fmt.Errorf("failed to insert monitor config: %w", err)
	}
	return nil
}

func AddHardDiskConfig(tx *sqlx.Tx, cfg models.Hard_disk_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO hard_disk_config (asset_id, type, storage)
		VALUES ($1, $2, $3)`,
		assetID, cfg.Type, cfg.Storage)
	if err != nil {
		return fmt.Errorf("failed to insert hard disk config: %w", err)
	}
	return nil
}

func AddPenDriveConfig(tx *sqlx.Tx, cfg models.Pen_drive_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO pendrive_config (asset_id, version, storage)
		VALUES ($1, $2, $3)`,
		assetID, cfg.Version, cfg.Storage)
	if err != nil {
		return fmt.Errorf("failed to insert pen drive config: %w", err)
	}
	return nil
}

func AddMobileConfig(tx *sqlx.Tx, cfg models.Mobile_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO mobile_config (asset_id, processor, ram, os, imei_1, imei_2)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, assetID, cfg.Processor, cfg.Ram, cfg.Os, cfg.IMEI1, cfg.IMEI2)
	if err != nil {
		return fmt.Errorf("failed to insert mobile config: %w", err)
	}
	return nil
}

func AddSimConfig(tx *sqlx.Tx, cfg models.Sim_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO sim_config (asset_id, number)
		VALUES ($1, $2)`,
		assetID, cfg.Number)
	if err != nil {
		return fmt.Errorf("failed to insert sim config: %w", err)
	}
	return nil
}

func AddAccessoryConfig(tx *sqlx.Tx, cfg models.Accessories_config_req, assetID uuid.UUID) error {
	_, err := tx.Exec(`
		INSERT INTO accessories_config (asset_id, type, additional_info)
		VALUES ($1, $2, $3)`,
		assetID, cfg.Type, cfg.AdditionalInfo)
	if err != nil {
		return fmt.Errorf("failed to insert accessory config: %w", err)
	}
	return nil
}
