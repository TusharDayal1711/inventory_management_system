package handler

import (
	jsoniter "github.com/json-iterator/go"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
	"strings"
)

func GetEmployeesWithFilters(w http.ResponseWriter, r *http.Request) {
	_, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}
	role := roles[0]
	if role != "admin" && role != "employee_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset manager can assign assets")
		return
	}
	var filter models.EmployeeFilter
	filter.SearchText = r.URL.Query().Get("search")
	filter.IsSearchText = filter.SearchText != ""

	if val := r.URL.Query().Get("type"); val != "" {
		filter.Type = strings.Split(val, ",")
	}
	if val := r.URL.Query().Get("role"); val != "" {
		filter.Role = strings.Split(val, ",")
	}
	if val := r.URL.Query().Get("asset_status"); val != "" {
		filter.AssetStatus = strings.Split(val, ",")
	}

	limit, offset := utils.GetPageLimitAndOffset(r)
	filter.Limit = limit
	filter.Offset = offset

	employees, err := dbhelper.GetFilteredEmployeesWithAssets(filter)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch employee data")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"employees": employees,
	})
}
