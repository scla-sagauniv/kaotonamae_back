package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {
	server.GET("/user", getAllUsers)
	server.POST("/createUser/:userId", postUserById)
}
