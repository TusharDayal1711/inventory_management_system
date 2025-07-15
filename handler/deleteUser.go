package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "employee_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can delete users")
		return
	}

	var req models.DeleteUserReq
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid user ID format")
		return
	}

	err = dbhelper.DeleteUserByID(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to soft delete user")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user deleted successfully",
	})
}
