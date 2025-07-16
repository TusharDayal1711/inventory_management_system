package dbhelper

import (
	"fmt"
	"inventory_management_system/database"
	"inventory_management_system/models"
	"strings"
)

func UpdateEmployeeInfo(req models.UpdateEmployeeReq) error {
	query := `UPDATE users SET `
	args := []interface{}{}
	argPos := 1

	if req.Username != "" {
		query += fmt.Sprintf("username = $%d, ", argPos)
		args = append(args, req.Username)
		argPos++
	}
	if req.Email != "" {
		query += fmt.Sprintf("email = $%d, ", argPos)
		args = append(args, req.Email)
		argPos++
	}
	if req.ContactNo != "" {
		query += fmt.Sprintf("contact_no = $%d, ", argPos)
		args = append(args, req.ContactNo)
		argPos++
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d AND archived_at IS NULL", argPos)
	args = append(args, req.UserID)

	result, err := database.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no user found or nothing updated")
	}

	return nil
}
