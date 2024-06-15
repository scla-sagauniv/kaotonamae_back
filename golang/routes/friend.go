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

func postFriendAdd(context echo.Context) error {
	myUserId := context.Param("myUserId")
	friendUserId := context.Param("friendUserId")
	groupMember, err := models.PostFriendAdd(myUserId, friendUserId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "フレンドを追加できませんでした。")
	}
	return context.JSON(http.StatusOK, groupMember)
}

func deleteFriendDelete(context echo.Context) error {
	myUserId := context.Param("myUserId")
	friendUserId := context.Param("friendUserId")
	err := models.DeleteFrinedDelete(myUserId, friendUserId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "フレンドを削除できませんでした。")
	}
	return context.JSON(http.StatusOK, "status: 完了")
}
