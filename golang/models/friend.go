package models

import (
	"time"
)

type Friend struct {
	UserId     string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255);index"`
	FriendId   string    `json:"friendId" gorm:"column:friend_id;primaryKey;type:VARCHAR(255)"`
	FriendName string    `json:"friendName" gorm:"column:friend_name;primaryKey;type:VARCHAR(255)"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at"`
}
