package handler

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
	"strings"
)

func GetAllAssetsWithFilters(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	role := roles[0]
	if role != "admin" && role != "asset_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can assign assets")
		return
	}
	var filter models.AssetFilter
	filter.SearchText = r.URL.Query().Get("search")
	if filter.SearchText != "" {
		filter.IsSearchText = true
		filter.SearchText = "%" + filter.SearchText + "%"
	}

	if val := r.URL.Query().Get("status"); val != "" {
		filter.Status = strings.Split(val, ",")
	}
	if val := r.URL.Query().Get("owned_by"); val != "" {
		filter.OwnedBy = strings.Split(val, ",")
	}
	if val := r.URL.Query().Get("type"); val != "" {
		filter.Type = strings.Split(val, ",")
	}
	limit, offset := utils.GetPageLimitAndOffset(r)
	filter.Limit = limit
	filter.Offset = offset
	assets, err := dbhelper.SearchAssetsWithFilter(filter)
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "failed to fetch records...", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{

		"assets": assets,
	})
}
