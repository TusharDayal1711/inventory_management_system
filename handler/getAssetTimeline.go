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
	assetId := r.URL.Query().Get("asset_id")
	assetUUID, err := uuid.Parse(assetId)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset_id")
		return
	}
	timeline, err := dbhelper.GetAssetTimeline(assetUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch asset timeline")
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"timeline": timeline,
	})

}
