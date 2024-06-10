package models

import (
	"time"
)

type Group struct {
	UserId    string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"`
	GroupId   string    `json:"groupId" gorm:"column:group_id;primaryKey;type:VARCHAR(255)"`
	GroupName string    `json:"groupName" gorm:"column:group_name;type:VARCHAR(255);index"`
	Overview  string    `json:"overview" gorm:"column:overview;type:VARCHAR(255);index"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}
