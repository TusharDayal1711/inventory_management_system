package handler

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/utils"
	"net/http"
)

func GetUserTimeline(w http.ResponseWriter, r *http.Request) {
	_, _, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	userId := r.URL.Query().Get("user_id")
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "falied to parse user uuid")
		return
	}
	timeline, err := dbhelper.GetUserAssetTimeline(userUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch timeline")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  userUUID,
		"timeline": timeline,
	})
}
