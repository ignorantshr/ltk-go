package entity

import (
	"common/enum"
)

type SkillFunc func(p *Player)

// 技能接口
type ISkill interface {
	GetName() string
	GetEvent() enum.Event
	GetFunc() SkillFunc
}

type Skill struct {
	Name  string
	Event enum.Event
	Func  SkillFunc
}

func NewSkill(name string, event enum.Event, function SkillFunc) ISkill {
	return &Skill{
		Name:  name,
		Event: event,
		Func:  function,
	}
}

func (s *Skill) GetName() string {
	return s.Name
}

func (s *Skill) GetEvent() enum.Event {
	return s.Event
}

func (s *Skill) GetFunc() SkillFunc {
	return s.Func
}
