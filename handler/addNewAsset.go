package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
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

	err = validator.New().Struct(assetReq)
	if err != nil {
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

	assetId, err := dbhelper.AddAsset(tx, assetReq, userUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to insert asset")
		return
	}
	switch assetReq.Type {
	case "laptop":
		var laptopConfig models.Laptop_config_req
		if err := json.Unmarshal(assetReq.Config, &laptopConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid laptop config format")
			return
		}
		err = dbhelper.AddLaptopConfig(tx, laptopConfig, assetId)

	case "mouse":
		var mouseConfig models.Mouse_config_req
		if err := json.Unmarshal(assetReq.Config, &mouseConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid mouse config format")
			return
		}
		err = dbhelper.AddMouseConfig(tx, mouseConfig, assetId)

	case "monitor":
		var monitorConfig models.Monitor_config_req
		if err := json.Unmarshal(assetReq.Config, &monitorConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid monitor config format")
			return
		}
		err = dbhelper.AddMonitorConfig(tx, monitorConfig, assetId)

	case "hard_disk":
		var hddConfig models.Hard_disk_config_req
		if err := json.Unmarshal(assetReq.Config, &hddConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid hard disk config format")
			return
		}
		err = dbhelper.AddHardDiskConfig(tx, hddConfig, assetId)

	case "pen_drive":
		var penConfig models.Pen_drive_config_req
		if err := json.Unmarshal(assetReq.Config, &penConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid pen drive config format")
			return
		}
		err = dbhelper.AddPenDriveConfig(tx, penConfig, assetId)

	case "mobile":
		var mobileConfig models.Mobile_config_req
		if err := json.Unmarshal(assetReq.Config, &mobileConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid mobile config format")
			return
		}
		err = dbhelper.AddMobileConfig(tx, mobileConfig, assetId)

	case "sim":
		var simConfig models.Sim_config_req
		if err := json.Unmarshal(assetReq.Config, &simConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid sim config format")
			return
		}
		err = dbhelper.AddSimConfig(tx, simConfig, assetId)

	case "accessory":
		var accessoryConfig models.Accessories_config_req
		if err := json.Unmarshal(assetReq.Config, &accessoryConfig); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid accessory config format")
			return
		}
		err = dbhelper.AddAccessoryConfig(tx, accessoryConfig, assetId)

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
		"msg":   "Asset and configuration created successfully",
		"Asset": assetReq,
	})
}
