package warrior

type ZhangFei struct {
	Warrior
}

func NewZhangFei() *Warrior {
	return NewWarrior("张飞", 4, 4)
}
