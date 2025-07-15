package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"
)

func DeleteAsset(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can delete assets")
		return
	}

	assetIDStr := r.URL.Query().Get("asset_id")
	assetID, err := uuid.Parse(assetIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset id")
		return
	}

	err = dbhelper.DeleteAssetByID(assetID)
	if err != nil {
		if err.Error() == "asset currently assigned to a user" {
			utils.RespondError(w, http.StatusConflict, err, "asset is currently assigned, failed be deleted")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to soft delete asset")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "asset deleted successfully",
	})
}
