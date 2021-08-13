package ygo

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Deck a deck
type Deck struct {
	Name  string
	Main  []*Card
	Extra []*Card
	Side  []*Card
}

// Card a card
type Card struct {
	Name    string
	Version string
}

// NewDeck create a Deck
func NewDeck(filename string) (*Deck, error) {
	var deck Deck

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return &deck, fmt.Errorf("read file '%s' error: %v", filename, err)
	}
	s := string(f)

	parts := strings.Split(s, "!")
	for _, part := range parts {
		if part == "" {
			continue
		}
		lines := strings.Split(part, "\n")

		if lines[0] != "Main" && lines[0] != "Extra" && lines[0] != "Side" {
			deck.Name = lines[0]
			continue
		}

		var cards []*Card

		length := len(lines)
		for i := 1; i < length; i++ {
			if lines[i] == "" {
				continue
			}
			vals := strings.Split(lines[i], " ")
			if len(vals) < 3 {
				return &deck, fmt.Errorf("invalid deck file: '%v' lack of infos", vals)
			}

			var (
				name    string
				version string
				count   int
			)

			name = vals[0]
			valsLen := len(vals)
			for j := 1; j < valsLen; j++ {
				if j < valsLen-2 {
					name += " " + vals[j]
				}
				if j == valsLen-2 {
					version = vals[j]
				}
				if j == valsLen-1 {
					count, err = strconv.Atoi(vals[j])
					if err != nil {
						return &deck, fmt.Errorf("invalid deck file: invalid number '%s'", vals[j])
					}
				}
			}
			for j := 0; j < count; j++ {
				cards = append(cards, &Card{
					Name:    name,
					Version: version,
				})
			}
		}

		switch lines[0] {
		case "Main":
			deck.Main = cards
		case "Extra":
			deck.Extra = cards
		case "Side":
			deck.Side = cards
		}
	}

	return &deck, nil
}

// List display the deck
func (d *Deck) List() {
	fmt.Printf("Deck: %s\n", d.Name)

	fmt.Printf("Main: %d\n", len(d.Main))
	for _, card := range d.Main {
		fmt.Printf("	%s %s\n", card.Name, card.Version)
	}

	fmt.Printf("Extra: %d\n", len(d.Extra))
	for _, card := range d.Extra {
		fmt.Printf("	%s %s\n", card.Name, card.Version)
	}

	fmt.Printf("Side: %d\n", len(d.Side))
	for _, card := range d.Side {
		fmt.Printf("	%s %s\n", card.Name, card.Version)
	}
}

// CreateDefaultDeckFile create a default deck file
func CreateDefaultDeckFile() string {
	str := `Zefra
!Main
智天之神星龙 SER 1
智天之神星龙 UTR 1
智天之神星龙 UR 1
神数的神意 R 3
神数的神托 R 3
星球改造 SER 1
秘龙星-神数囚牛 R 2
宝龙星-神数负屃 SER 1
觉星辉士-神数蝇王 N 1
救影依-神数纳迦 N 1
恩底弥翁的仆从 NPR 3
魔力统辖 SR 3
圣创魔导王 恩底弥翁 UR 1
魔导兽 胡狼王 SER 1
念动力反射者 R 3
紧急瞬间移动 SER 2
爆裂兽 R 1
爆裂模式 R 1
龙之灵庙 SR 1
龙之溪谷 SER 1
舞台旋转 NR 1
愚蠢的埋葬 UR 1
霸王眷龙 暗黑亚龙 N 1
霸王门 零 R 1
亡龙之战栗-死欲龙 SER 1
音响战士 吉他手 N 1
音响战士 架子鼓 N 1
喷气同调士 NPR 1
幻兽机 猎户座飞狮 R 1
龙落亲 R 1
妖精传姬-白雪 SER 1
神数的星战 N 1
龙星的九支 R 1
灰流丽 SER 3
PSY骨架装备·γ NPR 2
PSY骨架驱动者 NPR 1
增殖的G SER 1
无限泡影 SER 1
墓穴的指名者 SER 2
抹杀之指名者 SER 3
金满而谦虚之壶 SER 2
!Extra
鲜花女男爵 SER 1
装弹枪管狞猛龙 SER 1
辉龙星-蚣蝮 UTR 1
邪龙星-睚眦 UTR 1
月华龙 黑蔷薇 UR 1
流星登龙 UR 1
方程式运动员 电光赛道名将 UR 1
花园蔷薇花神 R 1
源龙星-望天吼 SER 1
虹光之宣告者 R 1
召命之神弓-阿波罗萨 SER 1
幻兽机 曙光女神百头龙 SER 1
刚炼装勇士·银金公主 SR 1
水晶机巧-继承玻纤 SER 1
警卫 路障山巨人 SER 1
!Side
恐龙摔跤手·潘克拉辛角龙 SR 1
古遗物-圣枪 R 3
增殖的G N 2
宇宙旋风 N 1
禁忌的一滴 SR 3
双龙卷 N 1
颉颃胜负 NPR 3
红色重启 N 1`
	ioutil.WriteFile("./zefra.deck", []byte(str), 0666)
	return "./zefra.deck"
}

// DeleteDefaultDeckFile delete the default deck file
func DeleteDefaultDeckFile() {
	os.Remove("./zefra.deck")
}
