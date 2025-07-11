package dbhelper

import (
	"fmt"
	"inventory_management_system/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateNewEmployee(tx *sqlx.Tx, req models.ManagerRegisterReq, managerUUID uuid.UUID) (uuid.UUID, error) {
	var userID uuid.UUID
	err := tx.Get(&userID, `
		INSERT INTO users (username, email, contact_no, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, req.Username, req.Email, req.ContactNo, managerUUID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert employee: %w", err)
	}

	// insert employee type in user_type table
	_, err = tx.Exec(`
		INSERT INTO user_type (user_id, type, created_by)
		VALUES ($1, $2, $3)
	`, userID, req.Type, managerUUID)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert employee type: %w", err)
	}
	return userID, nil
}
