package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {
	server.GET("/users", getAllUsers)
	server.POST("/createUser/:userId", postUserById)

	server.GET("/groups", getAllGroups)
	server.GET("/groups/:userId", getGroupsByUserId)
}
