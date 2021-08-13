package main

import (
	"fmt"
	"os"

	"ygo-calculator/ygo"
)

const (
	defaultDeckPath string = "./data/deck"
	defaultNeedPath string = "./data/need"
)

func main() {
	var (
		deckPath string
		needPath string
	)
	if len(os.Args) == 3 {
		deckPath = os.Args[1]
		needPath = os.Args[2]
	}
	if len(os.Args) != 3 {
		deckPath = defaultDeckPath
		needPath = defaultNeedPath
	}

	deck, err := ygo.NewDeck(deckPath)
	if err != nil {
		fmt.Printf("Create deck failed (%v). Use default deck.\n", err)
		path := ygo.CreateDefaultDeckFile()
		deck, _ = ygo.NewDeck(path)
		ygo.DeleteDefaultDeckFile()
	}
	need, err := ygo.NewCardsNeeded(needPath)
	if err != nil {
		fmt.Printf("Create need failed (%v). Use default need.\n", err)
		path := ygo.CreateDefaultNeedFile()
		need, _ = ygo.NewCardsNeeded(path)
		ygo.DeleteDefaultNeedFile()
	}

	deck.List()
	need.List()

	calculator := ygo.NewCalculator(deck)
	calculator.NeedCards(need.Cards, 100000)
}
