package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
)

func AddNewAssetWithConfig(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized user")
		return
	}

	var assetReq models.AddAssetWithConfigReq
	if err := utils.ParseJSONBody(r, &assetReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset input")
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

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid user ID")
		return
	}

	assetID, err := dbhelper.AddAsset(tx, assetReq, userUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to insert asset")
		return
	}

	switch assetReq.Type {
	case "laptop":
		var cfg models.Laptop_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid laptop config format")
			return
		}
		err = dbhelper.AddLaptopConfig(tx, cfg, assetID)

	case "mouse":
		var cfg models.Mouse_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid mouse config format")
			return
		}
		err = dbhelper.AddMouseConfig(tx, cfg, assetID)

	case "monitor":
		var cfg models.Monitor_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid monitor config format")
			return
		}
		err = dbhelper.AddMonitorConfig(tx, cfg, assetID)

	case "hard_disk":
		var cfg models.Hard_disk_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid hard disk config format")
			return
		}
		err = dbhelper.AddHardDiskConfig(tx, cfg, assetID)

	case "pen_drive":
		var cfg models.Pen_drive_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid pen drive config format")
			return
		}
		err = dbhelper.AddPenDriveConfig(tx, cfg, assetID)

	case "mobile":
		var cfg models.Mobile_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid mobile config format")
			return
		}
		err = dbhelper.AddMobileConfig(tx, cfg, assetID)

	case "sim":
		var cfg models.Sim_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid sim config format")
			return
		}
		err = dbhelper.AddSimConfig(tx, cfg, assetID)

	case "accessory":
		var cfg models.Accessories_config_req
		if err := json.Unmarshal(assetReq.Config, &cfg); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid accessory config format")
			return
		}
		err = dbhelper.AddAccessoryConfig(tx, cfg, assetID)

	default:
		utils.RespondError(w, http.StatusBadRequest, nil, "unsupported asset type")
		return
	}

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to insert asset config")
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "asset and config created successfully",
		"asset_id": assetID,
	})
}
