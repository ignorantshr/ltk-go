package entity

import (
	"common/enum"
)

func PaoXiao() ISkill {
	return NewSkill("咆哮", enum.BeginStageEvent, func(p *Player) {
		p.KillNumber = -1
	})
}
