package handler

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var userInfo models.PublicUserReq
	if err := utils.ParseJSONBody(r, &userInfo); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}
	err := validator.New().Struct(userInfo)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid asset input, email required")
		return
	}
	userID, err := dbhelper.GetUserByEmail(userInfo.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondError(w, http.StatusBadRequest, nil, "invalid email")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch user id")
		}
		return
	}
	fmt.Println("user ID ::", userID)
	userRole, err := dbhelper.GetUserRoleById(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondError(w, http.StatusBadRequest, nil, "invalid email")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch user")
		}
		return
	}
	fmt.Println("user role", userRole)

	accessToken, err := middlewares.GenerateJWT(userID.String(), []string{userRole})
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate access token")
		return
	}

	refreshToken, err := middlewares.GenerateRefreshToken(userID.String())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate refresh token")
		return
	}
	w.WriteHeader(http.StatusCreated)
	jsoniter.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "login success",
		"user id":       userID,
		"user role":     userRole,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
} //login
