package main

type Game struct {
	players []*Player
}

func newGame() *Game {
	return &Game{
		players: nil,
	}
}

func (g *Game) addTerrorlist(dressType string) {
	p := newPlayer("Terrorlist", dressType)
	g.players = append(g.players, p)
}

func (g *Game) addCounterTerrorlist(dressType string) {
	p := newPlayer("CounterTerrorlist", dressType)
	g.players = append(g.players, p)
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}
