package models

import (
	"kaotonamae_back/db"
	"time"

	"github.com/labstack/echo/v4"
)

type GroupMember struct {
	GroupId   string    `json:"groupId" gorm:"column:group_id;primaryKey;type:VARCHAR(255)"`
	UserId    string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"` // 所属メンバーのユーザーID
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

// 全グループメンバー取得処理
func GetAllGroupMembers() ([]GroupMember, error) {
	groupMembers := []GroupMember{}
	if db.DB.Find(&groupMembers).Error != nil {
		return nil, echo.ErrNotFound
	}
	return groupMembers, nil
}

// 特定のグループIDを付与されたグループメンバー取得処理
func GetGroupMembersByGroupId(id string) ([]GroupMember, error) {
	groupMembers := []GroupMember{}

	// groupIdがidであるグループメンバーを取得
	if err := db.DB.Where("group_id = ?", id).Find(&groupMembers).Error; err != nil {
		return nil, err
	}

	return groupMembers, nil
}
