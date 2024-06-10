package models

import (
	"kaotonamae_back/db"
	"time"

	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	UserId    string    `json:"userId" gorm:"column:user_id;type:VARCHAR(255);index"`
	UserName  string    `json:"userName" gorm:"column:user_name;type:VARCHAR(255)"`
	Furigana  string    `json:"furigana" gorm:"column:furigana;type:VARCHAR(255)"`
	Nickname  string    `json:"nickname" gorm:"column:nickname;type:VARCHAR(255)"`
	Gender    string    `json:"gender" gorm:"column:gender;type:VARCHAR(255)"`
	Photo     string    `json:"photo" gorm:"column:photo;type:VARCHAR(255)"`
	Birthday  string    `json:"birthday" gorm:"column:birthday;type:VARCHAR(255)"`
	Hobbys    string    `json:"hobbys" gorm:"column:hobbys;type:VARCHAR(255)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

// 全ユーザー取得処理
func GetAllUserInfos() ([]UserInfo, error) {
	userInfos := []UserInfo{}
	if db.DB.Find(&userInfos).Error != nil {
		return nil, echo.ErrNotFound
	}
	return userInfos, nil
}

// ユーザー取得処理(id)
func GetUserInfoById(id string) (*UserInfo, error) {
	userInfo := UserInfo{}
	if db.DB.Where("user_id = ?", id).First(&userInfo).Error != nil {
		return nil, echo.ErrNotFound
	}
	return &userInfo, nil
}
