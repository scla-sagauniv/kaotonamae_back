package models

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type quiz struct {
	QuizQuestionTop    string            `json:"quizQuestionTop"`
	QuizQuestionBottom string            `json:"quizQuestionBottom"`
	QuizAnswer         string            `json:"quizAnswer"`
	QuizHint           map[string]string `json:"quizHint"`
	UserPhoto          string            `json:"userPhoto"`
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

	for len(quizzes) < 15 {
		var randomMember GroupMember
		var userInfo *UserInfo

		// グループメンバーからランダムに1人を選ぶ
		attempts := 0
		for {
			if attempts >= len(groupMembers) {
				// すべてのメンバーがusedMembersの場合
				return quizzes, nil
			}
			randomMember = groupMembers[rand.Intn(len(groupMembers))]
			if !usedMembers[randomMember.UserId] {
				break
			}
			attempts++
		}

		// 選ばれたメンバーのユーザー情報を取得
		userInfo, err = GetUserInfoById(randomMember.UserId)
		if err != nil || userInfo == nil {
			continue
		}

		// クイズ候補となるフィールドを集める
		userInfoFields := map[string]string{
			"UserName":        userInfo.UserLastName + " " + userInfo.UserFirstName,
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
		hintCounter := 1
		for _, field := range validFields {
			if field != quizQuestion && hintCounter <= 3 {
				state := makeState(field)
				quizHint["Hint"+strconv.Itoa(hintCounter)] = state + userInfoFields[field] + "です。"
				hintCounter++
			}
		}

		// クイズを作成してリストに追加
		newQuiz := quiz{
			QuizQuestionTop:    "写真の人の",
			QuizQuestionBottom: "お名前は何でしょう",
			QuizAnswer:         quizAnswer,
			QuizHint:           quizHint,
			UserPhoto:          userInfo.Photo,
		}
		quizzes = append(quizzes, newQuiz)
		createdQuizzes[quizQuestion+quizAnswer] = true
		usedMembers[randomMember.UserId] = true

		// 名前以外のクイズ
		for len(quizzes) < 15 {
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

			quizField := remainingFields[rand.Intn(len(remainingFields))]
			var QuizQuestionBottomState string
			switch quizField {
			case "Nickname":
				QuizQuestionBottomState = "あだ名は何でしょう"
			case "Gender":
				QuizQuestionBottomState = "性別は何でしょう"
			case "Birthday":
				QuizQuestionBottomState = "お誕生日はいつでしょう"
			case "Age":
				QuizQuestionBottomState = "年齢はおいくつでしょう"
			case "Hobbys":
				QuizQuestionBottomState = "趣味は何でしょう"
			case "Organization":
				QuizQuestionBottomState = "所属は何でしょう"
			case "FavoriteColor":
				QuizQuestionBottomState = "好きな色は何色でしょう"
			case "FavoriteAnimal":
				QuizQuestionBottomState = "好きな動物は何でしょう"
			case "FavoritePlace":
				QuizQuestionBottomState = "好きな場所は何処でしょう"
			case "HolidayActivity":
				QuizQuestionBottomState = "休日の過ごし方は何でしょう"
			case "Weaknesses":
				QuizQuestionBottomState = "弱点は何でしょう"
			case "Language":
				QuizQuestionBottomState = "使う言語は何でしょう"
			default:
				QuizQuestionBottomState = "Default state"
			}
			quizAnswer = userInfoFields[quizField]

			// 既に生成されたクイズかどうかを確認
			if _, exists := createdQuizzes[quizQuestion+quizAnswer]; exists {
				continue
			}

			// ヒントを集める
			quizHint = make(map[string]string)
			hintCounter = 1
			for _, field := range validFields {
				if field != quizQuestion && hintCounter <= 3 {
					state := makeState(field)
					quizHint["Hint"+strconv.Itoa(hintCounter)] = state + userInfoFields[field] + "です。"
					hintCounter++
				}
			}

			// クイズを作成してリストに追加
			newQuiz = quiz{
				QuizQuestionTop:    userInfoFields["UserName"] + " さんの",
				QuizQuestionBottom: QuizQuestionBottomState,
				QuizAnswer:         quizAnswer,
				QuizHint:           quizHint,
				UserPhoto:          userInfo.Photo,
			}
			quizzes = append(quizzes, newQuiz)
			createdQuizzes[quizQuestion+quizAnswer] = true

			// 3/10の確率でループを抜ける
			if rand.Float32() < 0.25 {
				break
			}
		}
	}

	return quizzes, nil
}

func makeState(field string) string {
	var state string
	switch field {
	case "Nickname":
		state = "あだ名は "
	case "Gender":
		state = "性別は "
	case "Birthday":
		state = "お誕生日は "
	case "Age":
		state = "年齢は "
	case "Hobbys":
		state = "趣味は "
	case "Organization":
		state = "所属は "
	case "FavoriteColor":
		state = "好きな色は "
	case "FavoriteAnimal":
		state = "好きな動物は "
	case "FavoritePlace":
		state = "好きな場所は "
	case "HolidayActivity":
		state = "休日の過ごし方は "
	case "Weaknesses":
		state = "弱点は "
	case "Language":
		state = "使う言語は "
	default:
		state = "Default state"
	}
	return state
}
