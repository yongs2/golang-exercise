package poker

import (
	"io"
	"bufio"
	"strings"
	"fmt"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	fmt.Println("extractWinner.userInput=", userInput)
	return strings.Replace(userInput, " wins", "", 1)
}
