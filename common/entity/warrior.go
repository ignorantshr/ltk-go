package entity

/*
武将
 */

type IWarrior interface {
	DefaultSkills() []ISkill
}

type Warrior struct {
	Name string
	HP int	// 初始血量
	MaxHP int	// 初始血量上限
	Skills []ISkill
}

func NewWarrior(name string, HP, maxHP int) *Warrior {
	return &Warrior{
		Name: name,
		HP: HP,
		MaxHP: maxHP,
		Skills: []ISkill{},
	}
}

func (w *Warrior) DefaultSkills() []ISkill {
	return []ISkill{}
}

//func (w *Warrior) DispatchCards(cards []*entity.Card) {
//	w.HandCards = append(w.HandCards, cards...)
//}
//
//func (w *Warrior) AddHP(n int) {
//	w.HP += n
//}
//
//func (w *Warrior) SubHP(n int) {
//	w.HP -= n
//}
//
//func (w *Warrior) AfterDead() {
//	// drop all cards
//}
