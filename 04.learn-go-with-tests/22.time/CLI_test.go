package poker_test

import (
	"04.learn-go-with-tests/22.time"
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris wn from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(dummySpyAlerter, playerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris wins")
	})

	t.Run("record cleo wn from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(dummySpyAlerter, playerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo wins")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, playerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]

				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
