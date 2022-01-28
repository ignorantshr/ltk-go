package enum

/**
完整流程
 */

type IStage interface {
	getStageName() string
}

type Stage struct {
	Name string
}

func (s *Stage) getStageName() string {
	return s.Name
}

var BeginStage IStage
var JudgmentStage IStage
var DrawCardStage IStage
var PlayCardStage IStage
var DiscardCardStage IStage
var EndStage IStage

func init() {
	BeginStage = &Stage{
		Name: "准备阶段",
	}

	JudgmentStage = &Stage{
		Name: "判定阶段",
	}

	DrawCardStage = &Stage{
		Name: "摸牌阶段",
	}

	PlayCardStage = &Stage{
		Name: "出牌阶段",
	}

	DiscardCardStage = &Stage{
		Name: "弃牌阶段",
	}

	EndStage = &Stage{
		Name: "结束阶段",
	}
}

//BeginStage
//JudgmentStage
//DrawCardStage
//PlayCardStage
//DiscardCardStage
//EndStage

// 准备阶段
//type BeginStage struct {
//	Stage
//}

// 判定阶段
//type JudgmentStage struct {
//	Stage
//}

//// 摸牌阶段
//type DrawCardStage struct {
//	Stage
//}
//
//// 出牌阶段
//type PlayCardStage struct {
//	Stage
//}
//
//// 弃牌阶段
//type DiscardCardStage struct {
//	Stage
//}
//
//// 结束阶段
//type EndStage struct {
//	Stage
//}
