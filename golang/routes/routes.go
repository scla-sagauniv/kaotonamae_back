package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {
	server.GET("/users", getAllUsers)
	server.POST("/createUser/:userId", postUserById)

	server.GET("/auths", getAllAuths)
	server.POST("/auths/:userId/:email/:password", postCreateAuth)

	server.GET("/groups", getAllGroups)
	server.GET("/groups/:userId", getGroupsByUserId)
	server.GET("/group/:groupId", getGroupsByGroupId)
	server.POST("/newGroup/:userId", postNewGroup)
	server.PUT("/group", putGroupByGroupId)

	server.GET("/groupMembers", getAllGroupMembers)
	server.GET("/groupMembers/:groupId", getMembersByGroupId)
	server.POST("/groupMemberAdd/:groupId/:userId", postGroupMemberAdd)
	server.DELETE("/groupMemberDelete/:groupId/:userId", deleteGroupMemberDelete)

	server.GET("/friends", getAllFriends)
	server.GET("/friends/:userId", getFriendsById)
	server.POST("friendAdd/:myUserId/:friendUserId", postFriendAdd)
	server.DELETE("friendDelete/:myUserId/:friendUserId", deleteFriendDelete)
}
