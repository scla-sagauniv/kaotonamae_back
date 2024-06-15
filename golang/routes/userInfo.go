package routes

import (
	"kaotonamae_back/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllUserInfos(context echo.Context) error {
	userInfos, err := models.GetAllUserInfos()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザー情報を取得できませんでした。")
	}
	return context.JSON(http.StatusOK, userInfos)
}

func getUserInfoById(context echo.Context) error {
	userId := context.Param("userId")
	userInfo, err := models.GetUserInfoById(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザー情報を取得できませんでした。")
	}
	return context.JSON(http.StatusOK, userInfo)
}

func postCreateUserInfo(context echo.Context) error {
	userId := context.Param("userId")
	userInfo, err := models.PostCreateUserInfo(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザー情報を追加できませんでした。")
	}
	return context.JSON(http.StatusOK, userInfo)
}
