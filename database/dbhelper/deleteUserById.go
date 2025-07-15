package dbhelper

import (
	"fmt"
	"github.com/google/uuid"
	"inventory_management_system/database"
	"time"
)

func DeleteUserByID(userID uuid.UUID) error {
	tx, err := database.DB.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE users SET archived_at = $1
		WHERE id = $2 AND archived_at IS NULL
	`, now, userID)
	if err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	_, err = tx.Exec(`
		UPDATE user_roles SET archived_at = $1
		WHERE user_id = $2 AND archived_at IS NULL
	`, now, userID)
	if err != nil {
		return fmt.Errorf("failed to soft delete user roles: %w", err)
	}

	_, err = tx.Exec(`
		UPDATE user_type SET archived_at = $1
		WHERE user_id = $2 AND archived_at IS NULL
	`, now, userID)
	if err != nil {
		return fmt.Errorf("failed to soft delete user type: %w", err)
	}

	return nil
}
