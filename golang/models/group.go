package models

import (
	"kaotonamae_back/db"

	"time"

	"github.com/labstack/echo/v4"
)

type Group struct {
	UserId    string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"`
	GroupId   string    `json:"groupId" gorm:"column:group_id;primaryKey;type:VARCHAR(255)"`
	GroupName string    `json:"groupName" gorm:"column:group_name;type:VARCHAR(255);index"`
	Overview  string    `json:"overview" gorm:"column:overview;type:VARCHAR(255);index"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
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
func GetGroupByUserId(id string) ([]Group, error) {
	var groups []Group
	if err := db.DB.Where("user_id = ?", id).Find(&groups).Error; err == nil {
		return nil, nil
	}
	return groups, nil
}
