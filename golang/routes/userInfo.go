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

func putUserInfo(context echo.Context) error {
	var lateUserInfo models.UserInfo
	if err := context.Bind(&lateUserInfo); err != nil {
		return context.JSON(http.StatusBadRequest, "リクエストボディのバインドに失敗しました。")
	}
	// デバッグ用のログ出力
	if lateUserInfo.UserId == "" {
		return context.JSON(http.StatusBadRequest, "user_id が指定されていません。")
	}
	group, err := models.PutUserInfo(&lateUserInfo)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザー情報を更新できませんでした。")
	}
	return context.JSON(http.StatusOK, group)
}
