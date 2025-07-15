package handler

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"
)

func GetAssetTimeline(w http.ResponseWriter, r *http.Request) {
	_, _, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	assetIDStr := r.URL.Query().Get("asset_id")
	assetID, err := uuid.Parse(assetIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset id")
		return
	}

	timeline, err := dbhelper.GetAssetTimeline(assetID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch asset timeline")
		return
	}

	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"asset_id": assetID,
		"timeline": timeline,
	})

}
