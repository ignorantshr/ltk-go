package entity

import "math/rand"

/*
 * 牌堆
 */

type Pile struct {
	Cards []*Card
	DiscardPile []*Card
}

func InitPile(roomType RoomType) *Pile {
	pile := roomType.InitPile()
	pile.ShuffleCards()
	return pile
}

// DispatchCards 抽牌
func (pile *Pile) DispatchCards(num int, fromTop bool) []*Card {
	remainNum := len(pile.Cards)
	if fromTop {
		if remainNum > num {
			cards := pile.Cards[remainNum - num:]
			pile.Cards = pile.Cards[0:remainNum - num]
			return cards
		}else {
			cards := pile.Cards
			pile.ShuffleCards()
			cards = append(cards, pile.DispatchCards(num - remainNum, fromTop)...)
			return cards
		}
	}else {
		if remainNum >= num {
			cards := pile.Cards[:num]
			pile.Cards = pile.Cards[num:]
			return cards
		}else {
			cards := pile.Cards
			pile.ShuffleCards()
			cards = append(cards, pile.DispatchCards(num - remainNum, fromTop)...)
			return cards
		}
	}
}

// ShuffleCards 洗牌
func (pile *Pile) ShuffleCards() {
	var cards []*Card
	if len(pile.DiscardPile) > 0 {
		cards = pile.DiscardPile
	}else {
		cards = pile.Cards
	}
	for i := len(cards); i > 0; i-- {
		changeIndex := rand.Intn(i)
		cards[changeIndex], cards[i-1] = cards[i-1], cards[changeIndex]
	}
	pile.Cards = cards
	pile.DiscardPile = make([]*Card, 0)
}

// DiscardPatch 弃牌
func (pile *Pile) DiscardPatch(cards []*Card) {
	pile.DiscardPile = append(pile.DiscardPile, cards...)
}
