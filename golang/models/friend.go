package models

import (
	"kaotonamae_back/db"
	"time"

	"github.com/labstack/echo/v4"
)

type Friend struct {
	UserId     string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255);index"`
	FriendId   string    `json:"friendId" gorm:"column:friend_id;primaryKey;type:VARCHAR(255)"`
	FriendName string    `json:"friendName" gorm:"column:friend_name;primaryKey;type:VARCHAR(255)"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at"`
}

// 全フレンド取得処理
func GetAllFriends() ([]Friend, error) {
	friends := []Friend{}
	if db.DB.Find(&friends).Error != nil {
		return nil, echo.ErrNotFound
	}
	return friends, nil
}

// 特定のフレンド取得処理
func GetFriendsById(id string) ([]Friend, error) {
	friends := []Friend{}

	// userIdがidであるフレンドデータを取得
	if err := db.DB.Where("user_id = ?", id).Find(&friends).Error; err != nil {
		return nil, err
	}

	return friends, nil
}
