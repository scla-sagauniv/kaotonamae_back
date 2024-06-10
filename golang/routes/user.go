package routes

import (
	"kaotonamae_back/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllUsers(context echo.Context) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, users)
}

func getUserById(context echo.Context) error {
	userId := context.Param("userId")
	user, err := models.GetUserById(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, user)
}
