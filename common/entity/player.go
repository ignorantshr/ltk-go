package entity

import (
	"common/warrior"
)

type User struct {
	Id int64
	Name string
	Host string // IP address
	Port int
}

type Status int
const (
	Dead = iota
	Alive
)

type Player struct {
	User
	Status
	Role
	Index int
	HandCards []*Card
	Warrior warrior.IWarrior
}

func NewPlayer(u *User, index int, r Role, w warrior.IWarrior) *Player {
	p:= &Player{
		User: *u,
		Status: Alive,
		Index: index,
		Role: r,
		Warrior: w,
	}
	return p
}

func (p *Player) DispatchCards(cards []*Card) {
	p.HandCards = append(p.HandCards, cards...)
}
