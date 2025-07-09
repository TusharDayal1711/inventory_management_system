package dbhelper

import (
	"inventory_management_system/database"
	"inventory_management_system/models"
)

func AddNewAsset(assetReq models.AssetReq, userId string) error {
	_, err := database.DB.Exec(`
		INSERT INTO assets (
			brand, model, serial_no, purchase_date, 
		    owned_by, type, warranty_start, warranty_expire, 
		    added_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, assetReq.Brand, assetReq.Model, assetReq.SerialNo, assetReq.PurchaseDate,
		assetReq.OwnedBy, assetReq.Type, assetReq.WarrantyStart, assetReq.WarrantyExpire,
		userId)
	if err != nil {
		return err
	}
	return nil
}
