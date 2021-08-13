package ygo

import (
	"fmt"
	"math/rand"
	"time"
)

// Calculator a ygo calculator
type Calculator struct {
	Deck *Deck
}

var (
	minDeckSize int = 40
	initialDraw int = 5
)

// Init set the random seed
func Init() {
	rand.Seed(time.Now().Unix())
}

// NewCalculator create a calculator
func NewCalculator(deck *Deck) *Calculator {
	Init()
	return &Calculator{
		Deck: deck,
	}
}

// NeedCards 计算符合 needs 的抽卡数量出现的频率
// 注，needs 数组的规则：
//   1. needs[0] 表示第 1 张卡需要落在的范围，以此类推
//   2. 请把越精确的卡放在越前面，如：needs[0] = [A], needs[1] = [A, B, C]
//      如果反之，则若起手是 A 和 B，就会判定为不符合
//   3. 由于 2. 中说明的缘故，needs 中出现的卡最好不要重合。如果有特殊需要一定要重合，
//      则计算结果可能在某些特殊情况下会有一定程度的误差
func (c *Calculator) NeedCards(needs [][]*Card, drawTimes int) {
	var bingo int
	for i := 0; i < drawTimes; i++ {
		cards, err := c.ShuffleAndDraw(initialDraw)
		if err != nil {
			fmt.Printf("Darw failed (%v).\n", err)
			return
		}
		used := make([]bool, initialDraw)
		for index, cardsNeeded := range needs {
			found := false
			for order, card := range cards {
				for _, cardNeeded := range cardsNeeded {
					if !used[order] && (cardNeeded.Name == "" || cardNeeded.Name == card.Name) &&
						(cardNeeded.Version == "" || cardNeeded.Version == card.Version) {
						used[order], found = true, true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				break
			}
			if found && index == len(needs)-1 {
				bingo++
			}
		}
	}
	fmt.Printf("\nTotal: %d\nBingo: %d\n", drawTimes, bingo)
	fmt.Printf("Probability: %.2f%%\n\n", 100*float32(bingo)/float32(drawTimes))
}

// ShuffleAndDraw shuffle the deck and draw some cards
func (c *Calculator) ShuffleAndDraw(number int) ([]*Card, error) {
	var (
		shuffledDeck []*Card
		cardsDrawn   []*Card
	)

	if len(c.Deck.Main) < minDeckSize {
		return cardsDrawn, fmt.Errorf("invalid deck: cards less then 40")
	}
	if number <= 0 || number > len(c.Deck.Main) {
		return cardsDrawn, fmt.Errorf("invalid args")
	}

	shuffledDeck = append(shuffledDeck, c.Deck.Main...)

	for j := len(shuffledDeck) - 1; j > 0; j-- {
		num := rand.Intn(j + 1)
		shuffledDeck[j], shuffledDeck[num] = shuffledDeck[num], shuffledDeck[j]
	}

	for j := 0; j < number; j++ {
		cardsDrawn = append(cardsDrawn, shuffledDeck[j])
	}
	return cardsDrawn, nil
}
