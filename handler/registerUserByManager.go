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

func RegisterEmployeeByManager(w http.ResponseWriter, r *http.Request) {
	managerId, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	if len(roles) == 0 || (roles[0] != "admin" && roles[0] != "employee_manager") {
		utils.RespondError(w, http.StatusForbidden, fmt.Errorf("unauthorized role"), "only admin or employee_manager can register employees")
		return
	}

	var employeeReq models.ManagerRegisterReq
	if err := utils.ParseJSONBody(r, &employeeReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input body")
		return
	}

	if err := validator.New().Struct(employeeReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid role input")
		return
	}
	splitEmail := strings.Split(employeeReq.Email, "@")
	if len(splitEmail) != 2 || splitEmail[1] != "remotestate.com" {
		utils.RespondError(w, http.StatusBadRequest, nil, "only remotestate.com domain is valid")
		return
	}

	tx, err := database.DB.Beginx()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to start transaction")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	managerUUID, _ := uuid.Parse(managerId)
	newUserID, err := dbhelper.CreateNewEmployee(tx, employeeReq, managerUUID)
	if err != nil {
		tx.Rollback()
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create employee")
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"user created": newUserID,
	})
}
