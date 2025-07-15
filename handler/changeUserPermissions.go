package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
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

func ChangeUserRole(w http.ResponseWriter, r *http.Request) {
	adminID, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	adminUUID, err := uuid.Parse(adminID)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid admin ID")
		return
	}
	if len(roles) == 0 || roles[0] != "admin" {
		utils.RespondError(w, http.StatusForbidden, fmt.Errorf("unauthorized"), "only admin can update roles")
		return
	}
	var req models.UpdateUserRoleReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := validator.New().Struct(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid role input")
		return
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid user_id")
		return
	}
	tx, err := database.DB.Beginx()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "transaction failed")
		return
	}
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = dbhelper.UpdateUserRole(tx, userUUID, req.Role, adminUUID)
	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "already has the role") {
			utils.RespondError(w, http.StatusBadRequest, err, "user already has this role")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to update user role")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]string{
		"message": "user role changed successfully",
	})
}
