package handler

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	managerRole := roles[0]
	if managerRole != "admin" && managerRole != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can delete users")
		return
	}

	userID := r.URL.Query().Get("user_id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid user id...")
		return
	}

	userRole, err := dbhelper.GetUserRoleById(userUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch user role")
		return
	}

	if managerRole != "admin" && (userRole == "admin" || userRole == "asset_manager" || userRole == "inventory_manager") {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin can delete admin or manager roles")
		return
	}

	err = dbhelper.DeleteUserByID(userUUID)
	if err != nil {
		if err.Error() == "cannot delete user, still have asset assigned" {
			utils.RespondError(w, http.StatusConflict, err, "cannot delete user, still have asset assigned")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to delete user")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]string{
		"message": "user deleted successfully",
	})
}
