package handler

import (
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"

	"github.com/google/uuid"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can delete users")
		return
	}

	userID := r.URL.Query().Get("user_id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid user id...")
		return
	}

	err = dbhelper.DeleteUserByID(userUUID)
	if err != nil {
		if err.Error() == "cannot delete user, still have asset assigned" {
			utils.RespondError(w, http.StatusConflict, err, "cannot delete user, still have asset assigned")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to soft delete user")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]string{
		"message": "user soft deleted successfully",
	})
}
