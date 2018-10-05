package gocfr

// RhodeIslandGameState
type RhodeIslandGameState struct {
	round       Round
	parent      *RhodeIslandGameState
	causingMove Move
	table       PokerTable
	actors      map[ActionMakerIdentifier]ActionMaker
	nextToMove  ActionMakerIdentifier
	terminal    bool
}

func (state *RhodeIslandGameState) CurrentActor() ActionMaker {
	return state.actors[state.nextToMove]
}

func CreateRoot(playerAStack float64, playerBStack float64) RhodeIslandGameState {
	playerA := &PokerPlayer{id: PlayerA, availableMoves: nil, privateCards: []Card{}, stack: 100.0}
	playerB := &PokerPlayer{id: PlayerB, availableMoves: nil, privateCards: []Card{}, stack: 100.0}
	chance := &Chance{id: ChanceId, deck: CreateFullDeck()}

	actors := map[ActionMakerIdentifier]ActionMaker{PlayerA: playerA, PlayerB: playerB, ChanceId: chance}
	table := PokerTable{potSize: 0, publicCards: []Card{}}

	return RhodeIslandGameState{round: Start, table: table,
		actors: actors, nextToMove: ChanceId, causingMove: NoMove}
}

func (node *RhodeIslandGameState) CreateChild(round Round, move Move, nextToMove ActionMakerIdentifier, terminal bool) RhodeIslandGameState {
	child := RhodeIslandGameState{round: round,
		parent: node, causingMove: move, terminal: terminal,
		table: node.table, actors: node.actors, nextToMove: nextToMove}
	return child
}

func (state *RhodeIslandGameState) IsTerminal() bool {
	return state.terminal
}
