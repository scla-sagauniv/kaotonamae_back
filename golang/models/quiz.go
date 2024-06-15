package models

import (
	"errors"
	"math/rand"
	"time"
)

type quiz struct {
	QuizQuestion string   `json:"quizQuestion"`
	QuizAnswer   string   `json:"quizAnswer"`
	QuizHint     []string `json:"quizHint"`
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
			"UserLastName":      userInfo.UserLastName,
			"UserFirstName":     userInfo.UserFirstName,
			"LastNameFurigana":  userInfo.LastNameFurigana,
			"FirstNameFurigana": userInfo.FirstNameFurigana,
			"Nickname":          userInfo.Nickname,
			"Gender":            userInfo.Gender,
			"Birthday":          userInfo.Birthday,
			"Age":               userInfo.Age,
			"Hobbys":            userInfo.Hobbys,
			"Organization":      userInfo.Organization,
			"FavoriteColor":     userInfo.FavoriteColor,
			"FavoriteAnimal":    userInfo.FavoriteAnimal,
			"FavoritePlace":     userInfo.FavoritePlace,
			"HolidayActivity":   userInfo.HolidayActivity,
			"Weaknesses":        userInfo.Weaknesses,
			"Language":          userInfo.Language,
		}

		// ""でないフィールドを集める
		validFields := make(map[string]string)
		for field, value := range userInfoFields {
			if value != "" {
				validFields[field] = value
			}
		}

		if len(validFields) == 0 {
			continue
		}

		// クイズの質問と回答をランダムに選ぶ
		var quizQuestion string
		var quizAnswer string
		for question, answer := range validFields {
			quizQuestion = question
			quizAnswer = answer
			break
		}

		// 既に生成されたクイズかどうかを確認
		if _, exists := createdQuizzes[quizQuestion+quizAnswer]; exists {
			continue
		}

		// ヒントを集める
		var hints []string
		for question, answer := range validFields {
			if question != quizQuestion && len(hints) < 3 {
				hints = append(hints, answer)
			}
		}

		// クイズを作成してリストに追加
		newQuiz := quiz{
			QuizQuestion: quizQuestion,
			QuizAnswer:   quizAnswer,
			QuizHint:     hints,
		}
		quizzes = append(quizzes, newQuiz)
		createdQuizzes[quizQuestion+quizAnswer] = true
	}

	return quizzes, nil
}
