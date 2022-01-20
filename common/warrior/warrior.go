package warrior

/*
武将
 */

type IWarrior interface {
	AddHP(n int)
	SubHP(n int)
	AfterDead()
}

type Warrior struct {
	Name string
	HP int
	MaxHP int
	//HandCards []*entity.Card
}

func NewWarrior(name string, HP, maxHP int) *Warrior {
	return &Warrior{
		Name: name,
		HP: HP,
		MaxHP: maxHP,
	}
}

//func (w *Warrior) DispatchCards(cards []*entity.Card) {
//	w.HandCards = append(w.HandCards, cards...)
//}

func (w *Warrior) AddHP(n int) {
	w.HP += n
}

func (w *Warrior) SubHP(n int) {
	w.HP -= n
}

func (w *Warrior) AfterDead() {
	// drop all cards
}
