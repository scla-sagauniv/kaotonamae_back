package routes

import (
	"kaotonamae_back/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func getQuiz(context echo.Context) error {
	groupId := context.Param("groupId")
	quizzes, err := models.CreateQuizzesRess(groupId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "ユーザーを取得できませんでした。")
	}
	return context.JSON(http.StatusOK, quizzes)
}
