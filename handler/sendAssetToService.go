package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
)

func SendAssetToService(w http.ResponseWriter, r *http.Request) {
	managerId, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset are allowd")
		return
	}
	var req models.AssetServiceReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request payload")
		return
	}

	if err := validator.New().Struct(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid role input")
		return
	}
	managerUUID, err := uuid.Parse(managerId)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid manager UUID")
		return
	}
	if err := dbhelper.SendAssetForService(req, managerUUID); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "asset sent for servicing...",
	})
}
