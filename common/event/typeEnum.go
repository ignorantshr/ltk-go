package event

type Type int64

const (
	None Type = iota

	// 回合控制
	BeginStage // 开始阶段
	JudgmentStage // 判定阶段
	DrawCardStage // 摸牌阶段
	PlayCardStage // 出牌阶段
	DiscardCardStage // 弃牌阶段
	EndStage // 结束阶段

	Kill
	Flash
	Drink // 喝酒
	AddHP
	SubHP
)
