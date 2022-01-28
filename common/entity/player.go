package entity

import (
	"common/enum"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id int64
	Name string
	Host string // IP address
	Port int
	ResponseContext *gin.Context
}

type Status int
const (
	Dead = iota
	Alive
)

// 玩家
type Player struct {
	User
	Status                                   // 状态
	Role                                     // 角色
	Index       int                          // 位置
	HandCards   []*Card                      // 手牌
	Warrior     IWarrior                     // 武将
	Responsible map[enum.EventType]SkillFunc // 技能

	// 属性
	HP int	// 血量
	MaxHP int // 血量上限
	KillNumber int // 可出杀数量，-1 表示无限制
}

func NewPlayer(u *User, index int, r Role, w IWarrior) *Player {
	p:= &Player{
		User: *u,
		Status: Alive,
		Index: index,
		Role: r,
		Warrior: w,
		Responsible: make(map[enum.EventType]SkillFunc),
	}
	p.AddSkill(w.DefaultSkills()...)
	return p
}

func (p *Player) DispatchCards(cards []*Card) {
	p.HandCards = append(p.HandCards, cards...)
}

func (p *Player) AddSkill(skills... ISkill) {
	for _, skill := range skills {
		p.Responsible[skill.GetEvent().EventType] = skill.GetFunc()
	}
}
