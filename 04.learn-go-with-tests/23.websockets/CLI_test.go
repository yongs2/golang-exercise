package poker_test

import (
	"04.learn-go-with-tests/23.websockets"
	"bytes"
	"strings"
	"testing"
	"io"
	"net/http/httptest"
	"time"
	"github.com/gorilla/websocket"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled bool
	StartCalledWith int
	BlindAlert 	[]byte

	FinishedCalled bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayer int, out io.Writer) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayer
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishCalledWith = winner
}

const tenMS = 10 * time.Millisecond

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		poker.NewCLI(in, out, game).PlayPoker()

		assertMessagesSendToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("Pies")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSendToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("8", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSendToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})

	t.Run("start a game with 3 players, send some blind alerts down WS and declare Ruth the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "Ruth"

		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws" + strings.TrimPrefix(server.URL, "http") + "/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(tenMS)

		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
		within(t, tenMS, func() {
			assertWebsocketGotMsg(t, ws, wantedBlindAlert)
		})
	})
}

func assertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game shold not have finished")
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	passed := retryUntil(500 * time.Millisecond, func() bool {
		return game.FinishCalledWith == winner
	})

	if !passed {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertMessagesSendToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func within(t *testing.T, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <- time.After(d) :
		t.Error("timed out")
	case <- done:
	}
}

func assertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf(`got "%s", want "%s"`, string(msg), want)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
