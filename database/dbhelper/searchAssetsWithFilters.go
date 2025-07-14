package dbhelper

import (
	"fmt"
	"github.com/lib/pq"
	"inventory_management_system/database"
	"inventory_management_system/models"
)

func SearchAssetsWithFilter(filter models.AssetFilter) ([]models.AssetRes, error) {
	args := []interface{}{
		!filter.IsSearchText,
		filter.SearchText,
		pq.Array(filter.Status),
		pq.Array(filter.OwnedBy),
		pq.Array(filter.Type),
		filter.Limit,
		filter.Offset,
	}

	assetResponseModel := make([]models.AssetRes, 0)
	SQLquery := `SELECT id, 
       					brand, 
       					model, 
       					serial_no, 
       					type, 
       					owned_by, status 
			   			from assets  
			   			WHERE archived_at IS NULL
			   			AND ( $1 
			   			OR    brand     ILIKE $2
			            OR    model     ILIKE $2
			            OR    serial_no ILIKE $2
			       		   ) 
			        	AND status = ANY($3)
  						AND owned_by = ANY($4)
  						AND type = ANY($5)
				    	ORDER BY added_at DESC
						LIMIT $6 OFFSET $7`

	result, err := database.DB.Query(SQLquery, args...)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	fmt.Print(result)
	for result.Next() {
		var res models.AssetRes
		err := result.Scan(
			&res.ID, &res.Brand, &res.Model, &res.SerialNo,
			&res.Type, &res.OwnedBy, &res.Status,
		)
		if err != nil {
			return nil, err
		}
		assetResponseModel = append(assetResponseModel, res)
	}
	return assetResponseModel, nil

}
