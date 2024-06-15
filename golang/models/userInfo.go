package models

import (
	"errors"
	"kaotonamae_back/db"
	"time"

	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	UserId            string    `json:"userId" gorm:"column:user_id;primaryKey;type:VARCHAR(255);index"`
	UserLastName      string    `json:"userLastName" gorm:"column:user_last_name;type:VARCHAR(255)"`
	UserFirstName     string    `json:"userFirstName" gorm:"column:user_first_name;type:VARCHAR(255)"`
	LastNameFurigana  string    `json:"lastNameFurigana" gorm:"column:last_name_furigana;type:VARCHAR(255)"`
	FirstNameFurigana string    `json:"firstNameFurigana" gorm:"column:first_name_furigana;type:VARCHAR(255)"`
	Nickname          string    `json:"nickname" gorm:"column:nickname;type:VARCHAR(255)"` // あだ名
	Gender            string    `json:"gender" gorm:"column:gender;type:VARCHAR(255)"`     // 性別
	Photo             string    `json:"photo" gorm:"column:photo;type:VARCHAR(255)"`
	Birthday          string    `json:"birthday" gorm:"column:birthday;type:VARCHAR(255)"`                // 誕生日
	Age               string    `json:"age" gorm:"column:age;type:VARCHAR(255)"`                          // 年齢
	Hobbys            string    `json:"hobbys" gorm:"column:hobbys;type:VARCHAR(255)"`                    // 趣味
	Organization      string    `json:"organization" gorm:"column:organization;type:VARCHAR(255)"`        // 所属
	FavoriteColor     string    `json:"favoriteColor" gorm:"column:favorite_color;type:VARCHAR(255)"`     // 好きな色
	FavoriteAnimal    string    `json:"favoriteAnimal" gorm:"column:favorite_animal;type:VARCHAR(255)"`   // 好きな動物
	FavoritePlace     string    `json:"favoritePlace" gorm:"column:favorite_place;type:VARCHAR(255)"`     // 好きな場所
	HolidayActivity   string    `json:"holidayActivity" gorm:"column:holiday_activity;type:VARCHAR(255)"` // 休日の過ごし方
	Weaknesses        string    `json:"weaknesses" gorm:"column:weaknesses;type:VARCHAR(255)"`            // 弱点
	Language          string    `json:"language" gorm:"column:language;type:VARCHAR(255)"`                // 言語
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:created_at"`
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
	// Create new use
	userInfo := UserInfo{
		UserId:        id,
		UserLastName:  "User Last Name",
		UserFirstName: "User First Name",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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
	existingUserInfo.UserLastName = lateUserInfo.UserLastName
	existingUserInfo.UserFirstName = lateUserInfo.UserFirstName
	existingUserInfo.LastNameFurigana = lateUserInfo.LastNameFurigana
	existingUserInfo.FirstNameFurigana = lateUserInfo.FirstNameFurigana
	existingUserInfo.Nickname = lateUserInfo.Nickname
	existingUserInfo.Gender = lateUserInfo.Gender
	existingUserInfo.Photo = lateUserInfo.Photo
	existingUserInfo.Birthday = lateUserInfo.Birthday
	existingUserInfo.Hobbys = lateUserInfo.Hobbys
	existingUserInfo.Organization = lateUserInfo.Organization
	existingUserInfo.FavoriteColor = lateUserInfo.FavoriteColor
	existingUserInfo.FavoriteAnimal = lateUserInfo.FavoriteAnimal
	existingUserInfo.FavoritePlace = lateUserInfo.FavoritePlace
	existingUserInfo.HolidayActivity = lateUserInfo.HolidayActivity
	existingUserInfo.Weaknesses = lateUserInfo.Weaknesses

	existingUserInfo.UpdatedAt = lateUserInfo.UpdatedAt

	// データベースのレコードを更新
	if err := db.DB.Save(&existingUserInfo).Error; err != nil {
		return nil, errors.New("ユーザー更新中にエラーが発生しました")
	}

	return &existingUserInfo, nil
}
