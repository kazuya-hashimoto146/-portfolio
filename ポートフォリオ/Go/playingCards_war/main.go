package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	//初期化
	rand.Seed(time.Now().UnixNano())

	//カードを作成
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	deck := createDeck(suits, ranks)

	//プレイヤー名を入力
	var playerName string
	fmt.Println("あなたの名前を教えてください。")
	fmt.Scanln(&playerName)

	//何ゲームやるかを指定
	var playGameCount int
	for true {
		fmt.Println("ゲームのプレイ数を数値で指定してください。")
		_, err := fmt.Scanln(&playGameCount)

		if err != nil {
			fmt.Println("正しい数値ではありません。もう一度入力してください。")
			continue
		}
		break
	}
	fmt.Println("プレイ数：", playGameCount)

	//ゲーム開始
	fmt.Printf("戦争ゲームへようこそ %s さん！\n", playerName)
	fmt.Println("ゲームを開始します。")

	playerScore := 0
	computerScore := 0

	for i := 0; i < playGameCount; i++ {
		//プレイヤーがカードを引く
		var playerCard string
		fmt.Println("Enterキーを押してカードを引いてください。")
		fmt.Scanln()
		playerCard, deck = drawCard(deck)
		fmt.Printf("%s さんは %s を引きました\n", playerName, playerCard)

		//コンピューターがカードを引く
		var computerCard string
		computerCard, deck = drawCard(deck)
		fmt.Printf("CPUは %s を引きました\n", computerCard)

		//カードの強さを比較して勝敗を決定
		playerRank, _ := getRank(playerCard)
		computerRank, _ := getRank(computerCard)

		if playerRank > computerRank {
			fmt.Printf("%s さんの勝ちです！\n", playerName)
			playerScore += 3
		} else if playerRank < computerRank {
			fmt.Printf("CPUの勝ちです。")
			computerScore += 3
		} else {
			fmt.Println("引き分けです。次の戦争へ！")
			playerScore, computerScore = playWar(deck, playerScore, computerScore, playerName)
		}

		fmt.Printf("%s：%d、CPU：%d\n", playerName, playerScore, computerScore)
	}

	//ゲーム終了
	fmt.Println("ゲーム終了。")
	if playerScore > computerScore {
		fmt.Printf("%s さんの優勝！\n", playerName)
	} else if playerScore < computerScore {
		fmt.Println("CPUの優勝。")
	} else {
		fmt.Println("引き分け")
	}
}

// カードを作成する関数
func createDeck(suits []string, ranks []string) []string {
	deck := []string{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, suit+"の"+rank)
		}
	}
	return deck

}

// カードを1枚引く関数
func drawCard(deck []string) (string, []string) {
	cardIndex := rand.Intn(len(deck))
	card := deck[cardIndex]
	deck = append(deck[:cardIndex], deck[cardIndex+1:]...)
	return card, deck

}

// カードの強さ（ランク）を取得する関数
func getRank(card string) (int, error) {
	rank := card[strings.Index(card, "の")+3:]
	switch rank {
	case "Ace":
		return 14, nil
	case "King":
		return 13, nil
	case "Queen":
		return 12, nil
	case "Jack":
		return 11, nil
	default:
		return strconv.Atoi(rank)
	}

}

// カードのランクが同じ場合に行われる戦争を行う関数
func playWar(deck []string, playerScore int, computerScore int, playerName string) (int, int) {
	//カードを3枚引く
	playerCards := []string{}
	computerCards := []string{}
	for i := 0; i < 3; i++ {
		var playerCard string
		playerCard, deck = drawCard(deck)
		playerCards = append(playerCards, playerCard)

		var computerCard string
		computerCard, deck = drawCard(deck)
		computerCards = append(computerCards, computerCard)
	}

	//最後に引いたカードを比較する
	playlerLastCard := playerCards[len(playerCards)-1]
	computerLastCard := computerCards[len(computerCards)-1]

	playerRank, _ := getRank(playlerLastCard)
	computerRank, _ := getRank(computerLastCard)

	if playerRank > computerRank {
		fmt.Printf("%s さんの勝ちです！\n", playerName)
		playerScore += 3
	} else if playerRank < computerRank {
		fmt.Println("CPUの勝ちです。")
		computerScore += 3
	} else {
		fmt.Println("引き分けです。")

		//新しいデッキを作成して戦争を再開する
		var newDeck []string
		newDeck = append(newDeck, deck[:]...)
		var newPlayerScore, newComputerScore int
		newPlayerScore, newComputerScore = playWar(newDeck, playerScore, computerScore, playerName)
		playerScore += newPlayerScore
		computerScore += newComputerScore

	}

	return playerScore, computerScore

}
