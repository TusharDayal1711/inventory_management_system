package handler

import (
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
	"strings"
	"time"
)

func AddNewAsset(w http.ResponseWriter, r *http.Request) {
	userId, _, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized user")
		return
	}

	var newAssetModel models.AssetReq
	if err := utils.ParseJSONBody(r, &newAssetModel); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	// Validations
	if strings.TrimSpace(newAssetModel.Brand) == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "brand is required")
		return
	}
	if strings.TrimSpace(newAssetModel.Model) == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "model is required")
		return
	}
	if strings.TrimSpace(newAssetModel.SerialNo) == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "serial number is required")
		return
	}
	if newAssetModel.PurchaseDate.After(time.Now()) {
		utils.RespondError(w, http.StatusBadRequest, nil, "purchase date cannot be in the future")
		return
	}

	if !utils.IsAssetTypeValid(newAssetModel.Type) {
		utils.RespondError(w, http.StatusBadRequest, nil, "asset type is invalid")
		return
	}
	if newAssetModel.WarrantyStart.After(newAssetModel.WarrantyExpire) {
		utils.RespondError(w, http.StatusBadRequest, nil, "warranty start cannot be after expiry")
		return
	}

	err = dbhelper.AddNewAsset(newAssetModel, userId)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "conflict") {
			utils.RespondError(w, http.StatusConflict, err, "asset already exists")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, err, "failed to insert asset")
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]string{
		"message": "asset added successfully",
	})
}
