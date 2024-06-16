package models

import (
	"errors"
	"kaotonamae_back/db"
	"time"

	"github.com/labstack/echo/v4"
)

type GroupMember struct {
	GroupId     string    `json:"groupId" gorm:"column:group_id;primaryKey;type:VARCHAR(255)"`
	UserId      string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255)"` // 所属メンバーのユーザーID
	MemberName  string    `json:"memberName" gorm:"column:member_name;type:VARCHAR(255)"`
	MemberPhoto string    `json:"memberPhoto" gorm:"column:member_photo;type:VARCHAR(255)"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
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

// 特定のグループにグループメンバーを追加する
func PostGroupMemberAdd(userId, groupId string) (*GroupMember, error) {
	memberInfo, err := GetUserInfoById(userId)
	if err != nil {
		return nil, errors.New("フレンド情報の取得にエラーが発生しました")
	}
	memberFullName := memberInfo.UserLastName + " " + memberInfo.UserFirstName

	newGroupMember := GroupMember{
		GroupId:     groupId,
		UserId:      userId,
		MemberName:  memberFullName,
		MemberPhoto: memberInfo.Photo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.DB.Create(&newGroupMember).Error; err != nil {
		return nil, errors.New("グループメンバーの追加中にエラーが発生しました")
	}

	return &newGroupMember, nil
}

// 特定のグループからグループメンバーを削除する
func DeleteGroupMemberDelete(userId, groupId string) error {
	// 削除対象のGroupMemberを取得
	var groupMember GroupMember
	if err := db.DB.Where("user_id = ? AND group_id = ?", userId, groupId).First(&groupMember).Error; err != nil {
		return errors.New("グループメンバーが見つかりません")
	}

	// 削除
	if err := db.DB.Delete(&groupMember).Error; err != nil {
		return errors.New("グループメンバーの削除中にエラーが発生しました")
	}

	return nil
}
