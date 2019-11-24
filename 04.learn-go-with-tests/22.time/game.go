package poker

type Game interface {
	Start(numbferOfPlayers int)
	Finish(winner string)
}
