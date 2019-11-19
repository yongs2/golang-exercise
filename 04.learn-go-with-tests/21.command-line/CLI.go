package poker

import (

)

type CLI struct {
	playerStore PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}
