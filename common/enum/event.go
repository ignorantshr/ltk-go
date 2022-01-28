package enum

type IEvent interface {
	GetEvent() Event
}

type Event struct {
	Name string
	EventType
	SrcWarriorName string
	DestWarriorName string
}

func NewEvent(name string, eType EventType) Event {
	return Event{
		Name:      name,
		EventType: eType,
	}
}

type EventType int

const (
	NoneEvent EventType = iota

	// 回合控制
	BeginStageType // 开始阶段
	JudgmentStageType // 判定阶段
	//DrawCardStageType // 摸牌阶段
	//PlayCardStageType // 出牌阶段
	//DiscardCardStageType // 弃牌阶段
	//EndStageType // 结束阶段

	Kill
	Flash
	Drink // 喝酒
	AddHP
	SubHP
)

var BeginStageEvent Event
var JudgmentStageEvent Event

func init() {
	BeginStageEvent = NewBeginStageEvent()
	JudgmentStageEvent = NewJudgmentStageEvent()
}

func NewBeginStageEvent() Event {
	return NewEvent("开始阶段", BeginStageType)
}

func NewJudgmentStageEvent() Event {
	return NewEvent("判定阶段", JudgmentStageType)
}
