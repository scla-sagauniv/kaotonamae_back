package routes

import (
	"kaotonamae_back/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllGroupMembers(context echo.Context) error {
	groups, err := models.GetAllGroupMembers()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループメンバー取得できませんでした。")
	}
	return context.JSON(http.StatusOK, groups)
}

func getMembersByGroupId(context echo.Context) error {
	groupId := context.Param("groupId")
	groupListElements, err := models.GetGroupMembersByGroupId(groupId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループメンバーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, groupListElements)
}

func postGroupMemberAdd(context echo.Context) error {
	userId := context.Param("userId")
	groupId := context.Param("groupId")
	groupMember, err := models.PostGroupMemberAdd(userId, groupId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループメンバーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, groupMember)
}

func deleteGroupMemberDelete(context echo.Context) error {
	userId := context.Param("userId")
	groupId := context.Param("groupId")
	err := models.DeleteGroupMemberDelete(userId, groupId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループメンバーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, "status: 完了")
}
