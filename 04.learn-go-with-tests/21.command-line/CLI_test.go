package poker

import (
	"strings"
	"testing"
	"fmt"
)

func TestCLI(t *testing.T) {
	t.Run("record chris wn from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		fmt.Println("playerStore.chris.winCalls=", playerStore.winCalls)

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo wn from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		fmt.Println("playerStore.cleo.winCalls=", playerStore.winCalls)

		assertPlayerWin(t, playerStore, "Cleo")
	})
}
