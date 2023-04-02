package main

import "fmt"

func main() {
	totalScore := 0
	title(1, "楓", &totalScore)
	title(2, "ロビンソン", &totalScore)
	title(3, "チェリー", &totalScore)
	title(4, "空も飛べるはず", &totalScore)
	title(5, "春の歌", &totalScore)
	title(6, "スターゲイザー", &totalScore)
	title(7, "運命の人", &totalScore)

	fmt.Println("スコア", totalScore)
}

func title(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[問題%d] 次の楽曲タイトルを入力してください: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("正解！")
		*scorePtr += 10

	} else {
		fmt.Println("不正解…！")
	}
}
