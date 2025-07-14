package dbhelper

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func ReturnAsset(tx *sqlx.Tx, assetID, employeeID uuid.UUID, reason string) error {
	_, err := tx.Exec(`
		UPDATE asset_assign 
		SET returned_at = now(), return_reason = $1
		WHERE asset_id = $2 AND employee_id = $3 AND returned_at IS NULL AND archived_at IS NULL
	`, reason, assetID, employeeID)
	if err != nil {
		return fmt.Errorf("failed to update asset_assign: %w", err)
	}

	//updating asset table
	_, err = tx.Exec(`
		UPDATE assets SET status = 'available' WHERE id = $1
	`, assetID)
	if err != nil {
		return fmt.Errorf("failed to update asset status: %w", err)
	}
	return nil
}
