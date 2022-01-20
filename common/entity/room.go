package entity

type Room struct {
	Id int64
	Name string
	Players map[int]*Player
	Pile *Pile
}

type RoomType int
const (
	FiveRoom RoomType = iota
	EightRoom
	TwoRoom
)

func (r RoomType) PlayerNumber() int {
	switch r {
	default:
		fallthrough
	case FiveRoom:
		return 5
	case EightRoom:
		return 8
	case TwoRoom:
		return 2
	}
}

func (r RoomType) Roles() []Role {
	switch r {
	default:
		fallthrough
	case FiveRoom:
		roles := make([]Role, 3)
		roles[0] = Lord
		roles[1] = Traitor
		roles[2] = Loyalist
		for i:=0;i<2;i++ {
			roles = append(roles, Rebel)
		}
		return roles
	case EightRoom:
		roles := make([]Role, 2)
		roles[0] = Lord
		roles[1] = Traitor
		for i:=0;i<2;i++ {
			roles = append(roles, Loyalist)
		}
		for i:=0;i<4;i++ {
			roles = append(roles, Rebel)
		}
		return roles
	case TwoRoom:
		roles := make([]Role, 2)
		roles[0] = Lord
		roles[1] = Rebel
		return roles
	}
}

func (r RoomType) InitPile() *Pile {
	var cards []*Card
	switch r {
	default :
		for i := 0; i < 20; i++ {
			killCard := NewKillCard()
			cards = append(cards, &killCard)
		}
		for i := 0; i < 20; i++ {
			flashCard := NewFlashCard()
			cards = append(cards, &flashCard)
		}
		for i := 0; i < 5; i++ {
			liquorCard := NewLiquorCard()
			cards = append(cards, &liquorCard)
		}
		for i := 0; i < 5; i++ {
			peachCard := NewPeachCard()
			cards = append(cards, &peachCard)
		}
	}
	return &Pile{Cards: cards, DiscardPile: make([]*Card, 0)}
}
