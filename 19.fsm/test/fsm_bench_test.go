package fsm_bench

import (
	"fmt"
	"sync"
	"testing"

	cfsm "github.com/cocoonspace/fsm"
	lfsm "github.com/looplab/fsm"
)

// for cfsm
const (
	StateDisconnected cfsm.State = iota
	StateConnected
	StateUp
)

const (
	cEventConnect cfsm.Event = iota
	cEventDisconnect
	cEventEcho
)

// go tool pprof -png profile.out > cpu.png
// go tool pprof -png memprofile.out > mem.png
// gofmt -w -s *; go clean -testcache; go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out -run BenchmarkCocoonSpaceFSM
func BenchmarkCocoonSpaceFSM(b *testing.B) {
	f := cfsm.New(StateDisconnected)
	f.Transition(cfsm.On(cEventConnect), cfsm.Src(StateDisconnected), cfsm.Dst(StateConnected))
	f.Transition(cfsm.On(cEventDisconnect), cfsm.Src(StateConnected), cfsm.Dst(StateDisconnected))
	f.Transition(cfsm.On(cEventDisconnect), cfsm.Src(StateUp), cfsm.Dst(StateDisconnected))
	f.Transition(cfsm.On(cEventEcho), cfsm.Src(StateConnected), cfsm.Dst(StateUp))
	for i := 0; i < b.N; i++ {
		f.Event(cEventConnect)
		f.Event(cEventEcho)
		f.Event(cEventDisconnect)
	}
}

// for lfsm
const (
	EnterState = "enter_state"
	LeaveState = "leave_state"
	AfterEvent = "after_event"
)

// Link events
type LinkEventType int

const (
	EventConnect LinkEventType = iota
	EventDisconnect
	EventEcho
)

func (o LinkEventType) String() string {
	return [...]string{"connect", "disconnect", "echo"}[o]
}

func beforeHook(event LinkEventType) string {
	return fmt.Sprintf("before_%s", event)
}

func afterHook(event LinkEventType) string {
	return fmt.Sprintf("after_%s", event)
}

// Link states
var linkStatesOnce sync.Once

var linkStates *LinkStates

type LinkStates struct {
	Connected    string
	Disconnected string
	Up           string
	Any          []string // Any refers to all possible states
	Terminated   []string // Rejected, Killed, Failed, Completed
}

func NewLinkStates() *LinkStates {
	linkStatesOnce.Do(func() {
		linkStates = &LinkStates{
			Connected:    "Connected",
			Disconnected: "Disconnected",
			Up:           "Up",
			Any: []string{
				"Connected",
				"Disconnected",
				"Up",
			},
			Terminated: []string{
				"Disconnected",
			},
		}
	})
	return linkStates
}

// gofmt -w -s *; go clean -testcache; go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out -run BenchmarkLooplabFSM
func BenchmarkLooplabFSM(b *testing.B) {
	states := NewLinkStates()

	f := lfsm.NewFSM(
		states.Disconnected,
		lfsm.Events{
			{
				Name: EventConnect.String(),
				Src:  []string{states.Disconnected},
				Dst:  states.Connected,
			},
			{
				Name: EventDisconnect.String(),
				Src:  []string{states.Connected, states.Up},
				Dst:  states.Disconnected,
			},
			{
				Name: EventEcho.String(),
				Src:  []string{states.Connected},
				Dst:  states.Up,
			},
		},
		lfsm.Callbacks{},
	)
	for i := 0; i < b.N; i++ {
		f.Event(EventConnect.String())
		f.Event(EventEcho.String())
		f.Event(EventDisconnect.String())
	}
}
