package rhodeisland

import (
	"errors"
	. "github.com/int8/gopoker"
)

type Chance struct {
	id   ActorId
	deck Deck
}

func (chance *Chance) GetId() ActorId {
	return chance.id
}

type Player struct {
	Id      ActorId
	Card    *Card
	Stack   float32
	Actions []Action
}

func (player *Player) GetId() ActorId {
	return player.Id
}

func (player *Player) UpdateStack(stack float32) {
	player.Stack = stack
}

func (chance *Chance) Clone() *Chance {
	return &Chance{id: chance.id, deck: chance.deck.Clone()}
}

func (player *Player) Clone() *Player {
	return &Player{Card: player.Card, Id: player.Id, Stack: player.Stack, Actions: nil}
}

func (player *Player) Opponent() ActorId {
	return -player.Id
}

func (player *Player) CollectPrivateCard(card *Card) {
	player.Card = card
}

func (player *Player) PlaceBet(table *Table, betSize float32) {
	table.AddToPot(betSize)
	player.Stack -= betSize
}

func (player *Player) EvaluateHand(table *Table) []int8 {

	var flush, three, pair, straight, ownCard int8

	if (*player).Card.Suit == table.Cards[0].Suit && (*player).Card.Suit == table.Cards[1].Suit {
		flush = 1
	}

	if ((*player).Card.Name == table.Cards[0].Name) && ((*player).Card.Name == table.Cards[1].Name) {
		three = 1
	}

	if (((*player).Card.Name == table.Cards[0].Name) || ((*player).Card.Name == table.Cards[1].Name)) || table.Cards[0].Name == table.Cards[1].Name {
		pair = 1
	}

	if pair == 0 && cardsDiffersByTwo([]Card{*player.Card, table.Cards[0], table.Cards[1]}) {
		straight = 1
	}

	ownCard = int8((*player).Card.Name)

	return []int8{straight * flush, three, straight, flush, pair, ownCard}
}

func (player *Player) String() string {
	if player.Id == 1 {
		return "A"
	} else if player.Id == -1 {
		return "B"
	} else {
		return "Chance"
	}
	//TODO: not idiomatic !
	panic(errors.New("Code not reachable."))
}