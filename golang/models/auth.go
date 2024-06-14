package models

import (
	"errors"
	"kaotonamae_back/db"
	"time"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	UserId    string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255);index"`
	Email     string    `json:"email" gorm:"column:email;type:VARCHAR(255)"`
	Password  string    `json:"password" gorm:"column:password;type:VARCHAR(255)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

// 全グループ取得処理
func GetAllAuths() ([]Auth, error) {
	auths := []Auth{}
	if db.DB.Find(&auths).Error != nil {
		return nil, echo.ErrNotFound
	}
	return auths, nil
}

// 新規認証情追加処理
func PostCreateAuth(userId, email, password string) (*Auth, error) {
	newAuth := Auth{
		UserId:    userId,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.DB.Create(&newAuth).Error; err != nil {
		return nil, errors.New("認証情報の作成中にエラーが発生しました")
	}

	return &newAuth, nil
}
