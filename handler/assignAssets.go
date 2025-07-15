package handler

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
	"strings"
)

func AssignAssetToUser(w http.ResponseWriter, r *http.Request) {
	managerId, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can assign assets")
		return
	}

	var assignReq models.AssetAssignReq
	if err := utils.ParseJSONBody(r, &assignReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	userUUID, err := uuid.Parse(assignReq.UserID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid user id")
		return
	}
	assetUUID, err := uuid.Parse(assignReq.AssetID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset id")
		return
	}
	managerUUID, err := uuid.Parse(managerId)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid manager id")
		return
	}

	tx, err := database.DB.Beginx()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to start transaction")
		return
	}
	defer func() {
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = dbhelper.AssignAssetByID(tx, assetUUID, userUUID, managerUUID)
	if err != nil {
		if strings.Contains(err.Error(), "already assigned") {
			utils.RespondError(w, http.StatusConflict, err, "asset is already assigned")
			return
		}
		if strings.Contains(err.Error(), "failed to update assignment") {
			utils.RespondError(w, http.StatusInternalServerError, err, "failed to assign asset")
			return
		}
		if strings.Contains(err.Error(), "failed to insert into asset_assign table ") {
			utils.RespondError(w, http.StatusInternalServerError, err, "failed to insert into asset_assign table")
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "asset assigned successfully",
		"user_id":     userUUID,
		"asset_id":    assetUUID,
		"assigned_by": managerUUID,
	})
}
