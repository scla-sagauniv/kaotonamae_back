package routes

import (
	"kaotonamae_back/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllAuths(context echo.Context) error {
	auths, err := models.GetAllAuths()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "認証情報を取得できませんでした。")
	}
	return context.JSON(http.StatusOK, auths)
}

func postCreateAuth(context echo.Context) error {
	userId := context.Param("userId")
	email := context.Param("email")
	password := context.Param("password")
	auth, err := models.PostCreateAuth(userId, email, password)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "認証情報を追加できませんでした。")
	}
	return context.JSON(http.StatusOK, auth)
}
