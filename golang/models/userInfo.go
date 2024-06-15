package models

import (
	"errors"
	"kaotonamae_back/db"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	UserId          string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255);index"`
	UserName        string    `json:"userName" gorm:"column:user_name;type:VARCHAR(255)"`
	Furigana        string    `json:"furigana" gorm:"column:furigana;type:VARCHAR(255)"`
	Nickname        string    `json:"nickname" gorm:"column:nickname;type:VARCHAR(255)"`
	Gender          string    `json:"gender" gorm:"column:gender;type:VARCHAR(255)"`
	Photo           string    `json:"photo" gorm:"column:photo;type:VARCHAR(255)"`
	Birthday        string    `json:"birthday" gorm:"column:birthday;type:VARCHAR(255)"`
	Hobbys          string    `json:"hobbys" gorm:"column:hobbys;type:VARCHAR(255)"`
	Organization    string    `json:"organization" gorm:"column:organization;type:VARCHAR(255)"`
	HolidayActivity string    `json:"holidayActivity" gorm:"column:holiday_activity;type:VARCHAR(255)"`
	Weaknesses      string    `json:"weaknesses" gorm:"column:weaknesses;type:VARCHAR(255)"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt       time.Time `json:"createdAt" gorm:"column:created_at"`
}

// 全ユーザー情報取得処理
func GetAllUserInfos() ([]UserInfo, error) {
	userInfos := []UserInfo{}
	if db.DB.Find(&userInfos).Error != nil {
		return nil, echo.ErrNotFound
	}
	return userInfos, nil
}

// ユーザー情報取得処理(id)
func GetUserInfoById(id string) (*UserInfo, error) {
	userInfo := UserInfo{}
	if db.DB.Where("user_id = ?", id).First(&userInfo).Error != nil {
		return nil, echo.ErrNotFound
	}
	return &userInfo, nil
}

// ユーザー追加処理
func PostCreateUserInfo(id string) (*UserInfo, error) {
	nextUserName, err := findNextAvailableUserName(id)
	if err != nil {
		return nil, err
	}
	// Create new use
	userInfo := UserInfo{
		UserId:    id,
		UserName:  nextUserName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.DB.Create(&userInfo).Error; err != nil {
		return nil, errors.New("ユーザー作成中にエラーが発生しました")
	}

	return &userInfo, nil
}

// グループ情報の更新
func PutUserInfo(lateUserInfo *UserInfo) (*UserInfo, error) {
	// 現在のタイムスタンプを更新日時に設定
	lateUserInfo.UpdatedAt = time.Now()

	// データベースで該当のレコードを検索
	var existingUserInfo UserInfo
	if err := db.DB.Where("user_id = ?", lateUserInfo.UserId).First(&existingUserInfo).Error; err != nil {
		if err == echo.ErrNotFound {
			return nil, errors.New("ユーザーが見つかりません")
		}
		return nil, err
	}

	// 更新対象のフィールドを上書き
	existingUserInfo.UserName = lateUserInfo.UserName
	existingUserInfo.Furigana = lateUserInfo.Furigana
	existingUserInfo.Nickname = lateUserInfo.Furigana
	existingUserInfo.Gender = lateUserInfo.Gender
	existingUserInfo.Photo = lateUserInfo.Photo
	existingUserInfo.Birthday = lateUserInfo.Birthday
	existingUserInfo.Hobbys = lateUserInfo.Hobbys
	existingUserInfo.Organization = lateUserInfo.Organization
	existingUserInfo.HolidayActivity = lateUserInfo.HolidayActivity
	existingUserInfo.Weaknesses = lateUserInfo.Weaknesses

	existingUserInfo.UpdatedAt = lateUserInfo.UpdatedAt

	// データベースのレコードを更新
	if err := db.DB.Save(&existingUserInfo).Error; err != nil {
		return nil, errors.New("ユーザー更新中にエラーが発生しました")
	}

	return &existingUserInfo, nil
}

// 新規ユーザー追加の際にユーザーネームが重複しないようにする
func findNextAvailableUserName(userId string) (string, error) {
	var existingNames []string

	err := db.DB.Model(&UserInfo{}).
		Where("user_id = ?", userId).
		Pluck("user_name", &existingNames).Error
	if err != nil {
		return "", err
	}

	nextNumber := 1
	for {
		nextGroupName := "New User #" + strconv.Itoa(nextNumber)
		found := false
		for _, name := range existingNames {
			if name == nextGroupName {
				found = true
				break
			}
		}
		if !found {
			return nextGroupName, nil
		}
		nextNumber++
	}
}
