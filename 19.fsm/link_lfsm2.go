// Refer to https://gist.github.com/Champaneriyaraj/1d8e362a30f72f910c696d20b360fc5b
// Finite State Machine for link using https://github.com/looplab/fsm
//
// Refer
// https://github.com/apache/yunikorn-k8shim/blob/master/pkg/cache/node_state.go
// https://github.com/apache/yunikorn-k8shim/blob/master/pkg/cache/task_state.go
// https://github.com/dragonflyoss/Dragonfly2/blob/main/scheduler/resource/task.go
// https://github.com/dragonflyoss/Dragonfly2/blob/main/scheduler/resource/peer.go
// https://github.com/michelson/godard/blob/master/src/application/process.go
// https://github.com/opencord/voltha-epononu-adapter/blob/master/internal/pkg/onuadaptercore/uniportadmin.go
// https://github.com/choria-io/go-choria/blob/main/aagent/machine/machine.go
// https://github.com/instana/go-sensor/blob/master/fsm.go

package main

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/looplab/fsm"
)

// fsm
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

type Link struct {
	To           string
	EchoAttempts int
	FSM          *fsm.FSM
}

func NewLink(to string) *Link {
	l := &Link{
		EchoAttempts: 0,
	}

	states := NewLinkStates()

	l.FSM = fsm.NewFSM(
		states.Disconnected,
		fsm.Events{
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
		fsm.Callbacks{
			EnterState: func(e *fsm.Event) {
				l.enterState(e)
			},
			LeaveState: func(e *fsm.Event) {
				l.leaveState(e)
			},
			beforeHook(EventEcho): func(e *fsm.Event) {
				l.validateEchoAttempts(e)
			},
			afterHook(EventEcho): func(e *fsm.Event) {
				l.printEchoAttempts(e)
			},
			AfterEvent: func(e *fsm.Event) {
				l.afterEvent(e)
			},
		},
	)

	return l
}

func (l *Link) enterState(e *fsm.Event) {
	if e.FSM == nil {
		log.Printf("[Enter][%v] prev FSM is NIL", e.Event)
	} else {
		log.Printf("[Enter][%v][%v] The link was %s, now it is %s", e.Event, e.FSM.Current(), e.Src, e.Dst)
	}
}

func (l *Link) leaveState(e *fsm.Event) {
	if e.FSM == nil {
		log.Printf("[Leave][%v] prev FSM is NIL", e.Event)
	} else {
		log.Printf("[Leave][%v][%v], src[%v]->dst[%v]", e.Event, e.FSM.Current(), e.Src, e.Dst)
	}
}

func (l *Link) validateEchoAttempts(e *fsm.Event) {
	log.Println("Trying to ping")

	if l.EchoAttempts < 2 {
		l.EchoAttempts++
		e.Cancel(errors.New("Need At least 3 echo attempts"))
	}
}

func (l *Link) printEchoAttempts(e *fsm.Event) {
	log.Printf("EchoAttempts[%v]", l.EchoAttempts)
}

func (l *Link) afterEvent(e *fsm.Event) {
	log.Printf("[After][%v]", e.Event)
}

func (l *Link) Connect() {
	err := l.FSM.Event(EventConnect.String())
	if err != nil {
		log.Println(err)
	}
}

func (l *Link) Echo() {
	err := l.FSM.Event(EventEcho.String())
	if err != nil && err.Error() != "transition canceled with error: Need At least 3 echo attempts" {
		log.Println(err)
	}
}

func (l *Link) Disconnect() {
	err := l.FSM.Event(EventDisconnect.String())
	if err != nil {
		log.Println(err)
	}
}

func (l *Link) Visualize() {
	out, err := fsm.VisualizeWithType(l.FSM, fsm.GRAPHVIZ)
	if err != nil {
		log.Printf("VisualizeWithType.Err[%v]", err)
	} else {
		log.Printf("%v", out)
	}
}

func main() {
	link := NewLink(states.Disconnected)
	link.Connect()
	link.Echo()
	link.Echo()
	link.Echo()
	link.Disconnect()
	link.Visualize()
}
