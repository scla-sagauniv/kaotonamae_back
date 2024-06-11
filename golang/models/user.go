package models

import (
	"kaotonamae_back/db"

	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserId    string     `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"`
	UserInfos []UserInfo `gorm:"foreignKey:UserId"`
	Auths     []Auth     `gorm:"foreignKey:UserId"`
	Friends   []Friend   `gorm:"foreignKey:UserId"`
	Groups    []Group    `gorm:"foreignKey:UserId"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
}

// 全ユーザー取得処理
func GetAllUsers() ([]User, error) {
	users := []User{}
	if db.DB.Find(&users).Error != nil {
		return nil, echo.ErrNotFound
	}
	return users, nil
}

// ユーザー追加処理
func PostCreateUser(id string) (*User, error) {
	// Create new use
	user := User{
		UserId:    id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// if err := db.DB.Where("user_id = ?", id).First(&user).Error; err == nil {
	// 	return nil, nil
	// }
	if err := db.DB.Create(&user).Error; err != nil {
		return nil, errors.New("ユーザー作成中にエラーが発生しました")
	}

	return &user, nil
}
