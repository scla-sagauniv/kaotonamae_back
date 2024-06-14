package models

import (
	"errors"
	"kaotonamae_back/db"
	"strconv"

	"github.com/google/uuid"

	"time"

	"github.com/labstack/echo/v4"
)

type Group struct {
	UserId       string        `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"`
	GroupId      string        `json:"groupId" gorm:"column:group_id;primaryKey;type:VARCHAR(255)"`
	GroupName    string        `json:"groupName" gorm:"column:group_name;type:VARCHAR(255)"`
	Overview     string        `json:"overview" gorm:"column:overview;type:VARCHAR(255)"`
	GroupMembers []GroupMember `gorm:"foreignKey:groupId"`
	UpdatedAt    time.Time     `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt    time.Time     `json:"createdAt" gorm:"column:created_at"`
}

type GroupListElement struct {
	GroupId   string `json:"groupId" gorm:"column:group_id;primaryKey;type:VARCHAR(255)"`
	GroupName string `json:"groupName" gorm:"column:group_name;type:VARCHAR(255)"`
}

// 全グループ取得処理
func GetAllGroups() ([]Group, error) {
	groups := []Group{}
	if db.DB.Find(&groups).Error != nil {
		return nil, echo.ErrNotFound
	}
	return groups, nil
}

// 特定のグループ取得処理
func GetGroupByUserId(id string) ([]GroupListElement, error) {
	var groups []Group
	var groupListElements []GroupListElement

	// userIdがidであるグループデータを取得
	if err := db.DB.Where("user_id = ?", id).Find(&groups).Error; err != nil {
		return nil, err
	}

	// GroupListElementに整形して返す
	for _, group := range groups {
		groupListElement := GroupListElement{
			GroupId:   group.GroupId,
			GroupName: group.GroupName,
		}
		groupListElements = append(groupListElements, groupListElement)
	}

	return groupListElements, nil
}

// 新規グループ追加処理
func GetNewGroup(id string) (*Group, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	nextGroupName, err := findNextAvailableGroupName(id)
	if err != nil {
		return nil, err
	}

	// Create new group
	newGroup := Group{
		UserId:    id,
		GroupId:   uuid.String(),
		GroupName: nextGroupName,
		Overview:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.DB.Create(&newGroup).Error; err != nil {
		return nil, errors.New("グループ作成中にエラーが発生しました")
	}

	return &newGroup, nil
}

func findNextAvailableGroupName(userId string) (string, error) {
	var existingNames []string

	// Query existing group names for the given userId
	err := db.DB.Model(&Group{}).
		Where("user_id = ?", userId).
		Pluck("group_name", &existingNames).Error
	if err != nil {
		return "", err
	}

	// Find the next available number
	nextNumber := 1
	for {
		nextGroupName := "New Group #" + strconv.Itoa(nextNumber)
		found := false
		for _, name := range existingNames {
			if name == nextGroupName {
				found = true
				break
			}
		}
		if !found {
			return nextGroupName, nil
		}
		nextNumber++
	}
}
