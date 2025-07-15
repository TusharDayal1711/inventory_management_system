package dbhelper

import (
	"fmt"
	"github.com/google/uuid"
	"inventory_management_system/database"
	"inventory_management_system/models"
)

func GetAssetTimeline(assetUUID uuid.UUID) ([]models.AssetTimelineRes, error) {
	timeline := make([]models.AssetTimelineRes, 0)

	err := database.DB.Select(&timeline,
		`SELECT u.id as employee_id, 
       				  u.username as employee_name, 
       				  u.email,
       				  u.contact_no as contact_no,
       				  aa.assigned_at,
       				  aa.returned_at,
       				  aa.return_reason
				FROM users u 
				JOIN asset_assign aa ON u.id = aa.employee_id
				WHERE aa.asset_id = $1
				ORDER BY aa.assigned_at DESC;`, assetUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset timeline")
	}
	return timeline, nil
}
