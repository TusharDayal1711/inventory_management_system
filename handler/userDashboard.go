package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"
)

func GetUserDashboard(w http.ResponseWriter, r *http.Request) {
	userID, _, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid user id")
		return
	}

	dashboard, err := dbhelper.GetUserDashboardById(userUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch dashboard data")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dashboard)
}
