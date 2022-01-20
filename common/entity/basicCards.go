package entity

/**
杀,闪,酒,桃
Kill, flash, liquor, peach
 */

type KillCard struct {
	Card
}

func (c *KillCard) Clone() Card {
	return NewCard(
		c.Name,
		c.Type,
		c.Sort,
		c.Range,
	)
}

func NewKillCard() Card {
	return NewCard(Kill.String(), Kill, BasicCard, P2P)
}


type FlashCard struct {
	Card
}

func (c *FlashCard) Clone() Card {
	return NewCard(
		c.Name,
		c.Type,
		c.Sort,
		c.Range,
	)
}

func NewFlashCard() Card {
	return NewCard(Flash.String(), Flash, BasicCard, P2P)
}


type LiquorCard struct {
	Card
}

func (c *LiquorCard) Clone() Card {
	return NewCard(
		c.Name,
		c.Type,
		c.Sort,
		c.Range,
	)
}

func NewLiquorCard() Card {
	return NewCard(Liquor.String(), Liquor, BasicCard, P2P)
}


type PeachCard struct {
	Card
}

func (c *PeachCard) Clone() Card {
	return NewCard(
		c.Name,
		c.Type,
		c.Sort,
		c.Range,
	)
}

func NewPeachCard() Card {
	return NewCard(Peach.String(), Peach, BasicCard, P2P)
}
