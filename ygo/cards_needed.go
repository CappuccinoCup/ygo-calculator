package ygo

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// CardsNeeded cards needed
type CardsNeeded struct {
	Name  string
	Cards [][]*Card
}

// NewCardsNeeded create a CardsNeeded
func NewCardsNeeded(filename string) (*CardsNeeded, error) {
	var cardsNeeded CardsNeeded

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return &cardsNeeded, fmt.Errorf("read file '%s' error: %v", filename, err)
	}
	s := string(f)

	lines := strings.Split(s, "\n")

	cardsNeeded.Name = lines[0]
	count := len(lines)
	for i := 1; i < count; i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		parts := strings.Split(line, "+")

		var cards []*Card

		for _, part := range parts {
			if part == "" {
				continue
			}
			vals := strings.Split(part, " ")
			if len(vals) < 2 {
				return &cardsNeeded, fmt.Errorf("invalid need file: '%v' lack of infos", vals)
			}

			var (
				name    string
				version string
			)

			name = vals[0]
			valsLen := len(vals)
			for j := 1; j < valsLen; j++ {
				if j < valsLen-1 {
					name += " " + vals[j]
				}
				if j == valsLen-1 && vals[j] == "*" {
					version = ""
				}
				if j == valsLen-1 && vals[j] != "*" {
					version = vals[j]
				}
			}
			cards = append(cards, &Card{
				Name:    name,
				Version: version,
			})
		}

		if len(cards) > 0 {
			cardsNeeded.Cards = append(cardsNeeded.Cards, cards)
		}
	}

	if len(cardsNeeded.Cards) == 0 {
		return &cardsNeeded, fmt.Errorf("invalid need file: no cards needed")
	}
	return &cardsNeeded, nil
}

// List display the cards needed
func (n *CardsNeeded) List() {
	fmt.Printf("Need Cards: %s\n", n.Name)

	for index, cards := range n.Cards {
		fmt.Printf("%d: \n", index+1)
		for _, card := range cards {
			fmt.Printf("	%s %s\n", card.Name, card.Version)
		}
	}
}

// CreateDefaultNeedFile create a default need file
func CreateDefaultNeedFile() string {
	str := `Servant And Zefraath
智天之神星龙 *+神数的神意 *+神数的神托 *+星球改造 *+舞台旋转 *
恩底弥翁的仆从 *+魔力统辖 *`
	ioutil.WriteFile("Servant_Zefraath.need", []byte(str), 0666)
	return "./Servant_Zefraath.need"
}

// DeleteDefaultNeedFile delete the default need file
func DeleteDefaultNeedFile() {
	os.Remove("./Servant_Zefraath.need")
}
