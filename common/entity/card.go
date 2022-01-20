package entity

type RangeType int
const (
	P2P RangeType = iota	// one to one
	P2A	// one to all
	A2P	// all to one
)

type CardType int
const (
	Kill CardType = iota
	Flash
	Liquor
	Peach
)

func (ct CardType) String() string {
	switch ct {
	case Kill:
		return "杀"
	case Flash:
		return "闪"
	case Liquor:
		return "酒"
	case Peach:
		return "桃"
	default:
		return "-"
	}
}

type CardSort int
const (
	BasicCard CardSort = iota
	MagicCard
	EquipmentCard
)

type ICard interface {
	Clone() Card
}

type Card struct {
	Name string
	Type CardType
	Sort CardSort
	Range RangeType
}

func (c *Card) Clone() Card {
	return Card{
		c.Name,
		c.Type,
		c.Sort,
		c.Range,
	}
}

func NewCard(name string, cardType CardType, sort CardSort, rangeType RangeType) Card {
	return Card{
		name,
		cardType,
		sort,
		rangeType,
	}
}
