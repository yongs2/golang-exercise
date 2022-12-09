// Refer to https://gist.github.com/Champaneriyaraj/1d8e362a30f72f910c696d20b360fc5b
// Finite State Machine for link using https://github.com/looplab/fsm

package main

import (
	"errors"
	"log"

	"github.com/looplab/fsm"
)

type Link struct {
	To           string
	EchoAttempts int
	FSM          *fsm.FSM
}

func NewLink(to string) *Link {
	l := &Link{
		EchoAttempts: 0,
	}

	l.FSM = fsm.NewFSM(
		"disconnected",
		fsm.Events{
			{
				Name: "connect",
				Src:  []string{"disconnected"},
				Dst:  "connected",
			},
			{
				Name: "disconnect",
				Src:  []string{"connected", "up"},
				Dst:  "disconnected",
			},
			{
				Name: "echo",
				Src:  []string{"connected"},
				Dst:  "up",
			},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) {
				l.enterState(e)
			},
			"leave_state": func(e *fsm.Event) {
				l.leaveState(e)
			},
			"before_echo": func(e *fsm.Event) {
				l.validateEchoAttempts(e)
			},
			"after_echo": func(e *fsm.Event) {
				l.printEchoAttempts(e)
			},
			"after_event": func(e *fsm.Event) {
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
	err := l.FSM.Event("connect")
	if err != nil {
		log.Println(err)
	}
}

func (l *Link) Echo() {
	err := l.FSM.Event("echo")
	if err != nil && err.Error() != "transition canceled with error: Need At least 3 echo attempts" {
		log.Println(err)
	}
}

func (l *Link) Disconnect() {
	err := l.FSM.Event("disconnect")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	link := NewLink("disconnected")
	link.Connect()
	link.Echo()
	link.Echo()
	link.Echo()
	link.Disconnect()
}
