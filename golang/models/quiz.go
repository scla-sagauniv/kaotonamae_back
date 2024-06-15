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
	UserPhoto    string            `json:"userPhoto"`
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
	usedMembers := map[string]bool{}

	for len(quizzes) < 5 {
		var randomMember GroupMember
		var userInfo *UserInfo

		// グループメンバーからランダムに1人を選ぶ
		for {
			randomMember = groupMembers[rand.Intn(len(groupMembers))]
			if !usedMembers[randomMember.UserId] {
				break
			}
		}

		// 選ばれたメンバーのユーザー情報を取得
		userInfo, err = GetUserInfoById(randomMember.UserId)
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

		// 名前を問題にする
		quizQuestion := "UserName"
		quizAnswer := userInfoFields[quizQuestion]

		// 既に生成されたクイズかどうかを確認
		if _, exists := createdQuizzes[quizQuestion+quizAnswer]; exists {
			continue
		}

		// ヒントを集める
		quizHint := make(map[string]string)
		for _, field := range validFields {
			if field != quizQuestion && len(quizHint) < 3 {
				quizHint[field] = userInfoFields[field]
			}
		}

		// クイズを作成してリストに追加
		newQuiz := quiz{
			QuizQuestion: quizQuestion,
			QuizAnswer:   quizAnswer,
			QuizHint:     quizHint,
			UserPhoto:    userInfo.Photo,
		}
		quizzes = append(quizzes, newQuiz)
		createdQuizzes[quizQuestion+quizAnswer] = true
		usedMembers[randomMember.UserId] = true

		// 名前以外のクイズ
		for len(quizzes) < 5 {
			// 未使用のフィールドを集める
			var remainingFields []string
			for _, field := range validFields {
				if field != "UserName" && !createdQuizzes[field+userInfoFields[field]] {
					remainingFields = append(remainingFields, field)
				}
			}

			if len(remainingFields) == 0 {
				break
			}

			// そのメンバーの他の要素か、別のメンバーの名前を選ぶ
			if rand.Float32() < 0.5 {
				quizField := remainingFields[rand.Intn(len(remainingFields))]
				quizQuestion = quizField
				quizAnswer = userInfoFields[quizField]
			} else {
				for {
					randomMember = groupMembers[rand.Intn(len(groupMembers))]
					if !usedMembers[randomMember.UserId] {
						break
					}
				}
				userInfo, err = GetUserInfoById(randomMember.UserId)
				if err != nil || userInfo == nil {
					continue
				}
				quizQuestion = "UserName"
				quizAnswer = userInfo.UserLastName + " " + userInfo.UserFirstName
			}

			// 既に生成されたクイズかどうかを確認
			if _, exists := createdQuizzes[quizQuestion+quizAnswer]; exists {
				continue
			}

			// ヒントを集める
			quizHint = make(map[string]string)
			for _, field := range validFields {
				if field != quizQuestion && len(quizHint) < 3 {
					quizHint[field] = userInfoFields[field]
				}
			}

			// クイズを作成してリストに追加
			newQuiz = quiz{
				QuizQuestion: quizQuestion,
				QuizAnswer:   quizAnswer,
				QuizHint:     quizHint,
				UserPhoto:    userInfo.Photo,
			}
			quizzes = append(quizzes, newQuiz)
			createdQuizzes[quizQuestion+quizAnswer] = true

			// 1/2の確率でループを抜ける
			if rand.Float32() < 0.5 {
				break
			}
		}
	}

	return quizzes, nil
}
