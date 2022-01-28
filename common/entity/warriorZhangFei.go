package entity

type ZhangFei struct {
	Warrior
}

func NewZhangFei() *Warrior {
	return NewWarrior("张飞", 4, 4)
}

func (w *ZhangFei) DefaultSkills() []ISkill {
	return []ISkill{PaoXiao()}
}
