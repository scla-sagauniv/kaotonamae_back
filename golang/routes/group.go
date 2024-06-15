package routes

import (
	"kaotonamae_back/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllGroups(context echo.Context) error {
	groups, err := models.GetAllGroups()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, groups)
}

func getGroupsByUserId(context echo.Context) error {
	userId := context.Param("userId")
	groupListElements, err := models.GetGroupByUserId(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, groupListElements)
}

func postNewGroup(context echo.Context) error {
	userId := context.Param("userId")
	group, err := models.PostNewGroup(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループを追加できませんでした。")
	}
	return context.JSON(http.StatusOK, group)
}

func putGroupByGroupId(context echo.Context) error {
	var lateGroup models.Group
	if err := context.Bind(&lateGroup); err != nil {
		return context.JSON(http.StatusBadRequest, "リクエストボディのバインドに失敗しました。")
	}
	// デバッグ用のログ出力
	if lateGroup.GroupId == "" {
		return context.JSON(http.StatusBadRequest, "group_id が指定されていません。")
	}
	group, err := models.PutGroupByGroupId(&lateGroup)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループを更新できませんでした。")
	}
	return context.JSON(http.StatusOK, group)
}
