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

func RetrieveAsset(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin or asset manager can return assets")
		return
	}

	var req models.AssetReturnReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request")
		return
	}

	assetUUID, err := uuid.Parse(req.AssetID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset_id")
		return
	}

	employeeUUID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid employee_id")
		return
	}

	tx, err := database.DB.Beginx()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to begin transaction")
		return
	}
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = dbhelper.RetrieveAsset(tx, assetUUID, employeeUUID, req.ReturnReason)
	if err != nil {
		if strings.Contains(err.Error(), "no matching asset assignment found") {
			utils.RespondError(w, http.StatusNotFound, nil, " no such asset exist or already returned")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, err, "no matching asset assignment found")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]string{
		"message": "asset returned successfully...",
	})
}
