package handler

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"
)

func ReceivedFromService(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can mark asset as received")
		return
	}

	assetIDStr := r.URL.Query().Get("asset_id")
	assetUUID, err := uuid.Parse(assetIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset UUID")
		return
	}

	err = dbhelper.RecivedAssetFromService(assetUUID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Asset received",
		"asset_id": assetUUID.String(),
	})
}
