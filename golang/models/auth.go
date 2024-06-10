package models

import (
	"time"
)

type Auth struct {
	UserId    string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255);index"`
	Email     string    `json:"email" gorm:"column:email;type:VARCHAR(255)"`
	Password  string    `json:"password" gorm:"column:password;type:VARCHAR(255)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}
