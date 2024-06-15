package models

import (
	"errors"
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

// フレンドの登録をする
func PostFriendAdd(myUserId, friendUserId string) (*Friend, error) {
	// friendInfo, err := GetUserInfoById(friendUserId)
	// if err != nil {
	// 	return nil, errors.New("フレンド情報の取得にエラーが発生しました")
	// }

	newFrined := Friend{
		UserId:     myUserId,
		FriendId:   friendUserId,
		FriendName: "",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := db.DB.Create(&newFrined).Error; err != nil {
		return nil, errors.New("フレンドの登録中にエラーが発生しました")
	}

	return &newFrined, nil
}

// フレンドの削除をする
func DeleteFrinedDelete(myUserId, friendUserId string) error {
	var friend Friend
	if err := db.DB.Where("user_id = ? AND friend_id = ?", myUserId, friendUserId).First(&friend).Error; err != nil {
		return errors.New("該当のフレンドが見つかりません")
	}

	// 削除
	if err := db.DB.Delete(&friend).Error; err != nil {
		return errors.New("該当のフレンドの削除中にエラーが発生しました")
	}

	return nil
}
