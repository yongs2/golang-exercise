package poker

import (
	"io"
)

type Game interface {
	Start(numbferOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}
