package handler

import (
	"github.com/go-playground/validator/v10"
	"inventory_management_system/database"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

func PublicRegister(w http.ResponseWriter, r *http.Request) {
	var userReq models.PublicUserReq
	if err := utils.ParseJSONBody(r, &userReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	//validate using playground validator
	err := validator.New().Struct(userReq)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset input")
		return
	}
	splitEmail := strings.Split(userReq.Email, "@")
	if len(splitEmail) != 2 || splitEmail[1] != "remotestate.com" {
		utils.RespondError(w, http.StatusBadRequest, nil, "only remotestate.com domain is valid")
		return
	}
	usernameParts := strings.Split(splitEmail[0], ".")
	if len(usernameParts) != 2 {
		utils.RespondError(w, http.StatusBadRequest, nil, "failed to get username from email")
		return
	}
	username := usernameParts[0] + " " + usernameParts[1]

	//transaction
	tx, err := database.DB.Beginx()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to start transaction")
		return
	}

	var userId uuid.UUID
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// checking if user with same email already exit
	isExists, err := dbhelper.IsUserExists(tx, userReq.Email)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to check existing user")
		return
	}
	if isExists {
		utils.RespondError(w, http.StatusBadRequest, nil, "email already registered")
		return
	}

	//insertin new user in users table with username and email
	userId, err = dbhelper.InsertIntoUser(tx, username, userReq.Email)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}
	const defaultRole = "employee"
	const defaultType = "full_time"

	if err = dbhelper.InsertIntoUserRole(tx, userId, defaultRole, userId); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to assign role")
		return
	}

	if err = dbhelper.InsertIntoUserType(tx, userId, defaultType, userId); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to assign type")
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message": "account created successfully",
		"userId":  userId,
	})
}
