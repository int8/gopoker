package gocfr

import "fmt"

type Move int8
const (
	Check Move = iota
	Bet
	Raise
	Call
	Fold
	DealPublicCard
	DealPrivateCards
)

type Action struct {
	player Player
	move   Move
}

type ActionsCache struct {
	actions []Action
}

func (a Action) String() string {
	return fmt.Sprintf("%v:%v", a.player, a.move)
}

func (m Move) String() string {
	switch m {
	case Check:
		return "Check"
	case Bet:
		return "Bet"
	case Call:
		return "Call"
	case Fold:
		return "Fold"
	case Raise:
		return "Raise"
	case DealPrivateCards:
		return "DealPrivateCards"
	case DealPublicCard:
		return "DealPublicCard"
	}
	return "Undefined"
}
