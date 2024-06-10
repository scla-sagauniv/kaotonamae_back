package models

import (
	"time"
)

type User struct {
	UserId    string     `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"`
	UserInfos []UserInfo `gorm:"foreignKey:UserId"`
	Auths     []Auth     `gorm:"foreignKey:UserId"`
	Friends   []Friend   `gorm:"foreignKey:UserId"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
}
