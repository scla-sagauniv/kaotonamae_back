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

func GetNewGroup(context echo.Context) error {
	userId := context.Param("userId")
	group, err := models.GetNewGroup(userId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "グループを追加できませんでした。")
	}
	return context.JSON(http.StatusOK, group)
}
