package main

import "fmt"

func main() {
	totalScore := 0
	title(1, "make total destroy", &totalScore)
	title(2, "marigold", &totalScore)
	title(3, "letter experiment", &totalScore)
	title(4, "graveless", &totalScore)
	title(5, "alpha", &totalScore)
	title(6, "zero", &totalScore)
	title(7, "blood eagle", &totalScore)

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
