package models

import (
	"errors"
	"math/rand"
	"time"
)

type quiz struct {
	QuizQuestion string            `json:"quizQuestion"`
	QuizAnswer   string            `json:"quizAnswer"`
	QuizHint     map[string]string `json:"quizHint"`
}

func CreateQuizzesRess(GroupId string) ([]quiz, error) {
	var quizzes []quiz

	// GroupMembersを取得
	groupMembers, err := GetGroupMembersByGroupId(GroupId)
	if err != nil || len(groupMembers) == 0 {
		return nil, errors.New("グループメンバーが見つかりません")
	}

	// ランダムに選択するための乱数シードを設定
	rand.Seed(time.Now().UnixNano())

	createdQuizzes := map[string]bool{}

	for len(quizzes) < 5 {
		// グループメンバーからランダムに1人を選ぶ
		randomMember := groupMembers[rand.Intn(len(groupMembers))]

		// 選ばれたメンバーのユーザー情報を取得
		userInfo, err := GetUserInfoById(randomMember.UserId)
		if err != nil || userInfo == nil {
			continue
		}

		// クイズ候補となるフィールドを集める
		userInfoFields := map[string]string{
			"UserName":        userInfo.UserLastName + " " + userInfo.UserFirstName,
			"Furigana":        userInfo.LastNameFurigana + " " + userInfo.FirstNameFurigana,
			"Nickname":        userInfo.Nickname,
			"Gender":          userInfo.Gender,
			"Birthday":        userInfo.Birthday,
			"Age":             userInfo.Age,
			"Hobbys":          userInfo.Hobbys,
			"Organization":    userInfo.Organization,
			"FavoriteColor":   userInfo.FavoriteColor,
			"FavoriteAnimal":  userInfo.FavoriteAnimal,
			"FavoritePlace":   userInfo.FavoritePlace,
			"HolidayActivity": userInfo.HolidayActivity,
			"Weaknesses":      userInfo.Weaknesses,
			"Language":        userInfo.Language,
		}

		// ""でないフィールドを集める
		var validFields []string
		for field, value := range userInfoFields {
			if value != "" {
				validFields = append(validFields, field)
			}
		}

		if len(validFields) == 0 {
			continue
		}

		// クイズの質問と回答をランダムに選ぶ
		quizField := validFields[rand.Intn(len(validFields))]
		quizQuestion := quizField
		quizAnswer := userInfoFields[quizField]

		// 既に生成されたクイズかどうかを確認
		if _, exists := createdQuizzes[quizQuestion+quizAnswer]; exists {
			continue
		}

		// ヒントを集める
		quizHint := make(map[string]string)
		for _, field := range validFields {
			if field != quizField && len(quizHint) < 3 {
				quizHint[field] = userInfoFields[field]
			}
		}

		// クイズを作成してリストに追加
		newQuiz := quiz{
			QuizQuestion: quizQuestion,
			QuizAnswer:   quizAnswer,
			QuizHint:     quizHint,
		}
		quizzes = append(quizzes, newQuiz)
		createdQuizzes[quizQuestion+quizAnswer] = true
	}

	return quizzes, nil
}
