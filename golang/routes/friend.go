package routes

import (
	"kaotonamae_back/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllFriends(context echo.Context) error {
	friends, err := models.GetAllFriends()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "フレンドを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, friends)
}

func getFriendsById(context echo.Context) error {
	userId := context.Param("userId")
	friends, err := models.GetFriendsById(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "フレンドを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, friends)
}
