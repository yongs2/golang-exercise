// Refer to https://gist.github.com/Champaneriyaraj/1d8e362a30f72f910c696d20b360fc5b
// Finite State Machine for link using https://github.com/cocoonspace/fsm

package main

import (
	"log"

	"github.com/cocoonspace/fsm"
)

const (
	StateDisconnected fsm.State = iota
	StateConnected
	StateUp
)

func getStateString(state fsm.State) string {
	return [...]string{"Disconnected", "Connected", "Up"}[state]
}

const (
	EventConnect fsm.Event = iota
	EventDisconnect
	EventEcho
)

func getEventString(event fsm.Event) string {
	return [...]string{"Connect", "Disconnect", "Echo"}[event]
}

type Link struct {
	To           string
	EchoAttempts int
	FSM          *fsm.FSM
}

func NewLink(to fsm.State) *Link {
	l := &Link{
		EchoAttempts: 0,
	}

	l.FSM = fsm.New(to)
	l.FSM.Transition(
		fsm.On(EventConnect), fsm.Src(StateDisconnected), fsm.Dst(StateConnected),
	)
	l.FSM.Transition(
		fsm.On(EventDisconnect), fsm.Src(StateConnected), fsm.Dst(StateDisconnected),
	)
	l.FSM.Transition(
		fsm.On(EventDisconnect), fsm.Src(StateUp), fsm.Dst(StateDisconnected),
	)
	l.FSM.Transition(
		fsm.On(EventEcho), fsm.Src(StateConnected), fsm.Check(l.validateEchoAttempts), fsm.Dst(StateUp),
	)

	l.FSM.Enter(l.enterState)
	l.FSM.Exit(l.leaveState)

	return l
}

func (l *Link) enterState(state fsm.State) {
	log.Printf("[Enter][%v]", getStateString(state))
}

func (l *Link) leaveState(state fsm.State) {
	log.Printf("[Exit][%v]", getStateString(state))
}

func (l *Link) validateEchoAttempts() bool {
	log.Println("Trying to ping")
	if l.EchoAttempts < 2 {
		l.EchoAttempts++
		return false
	} else {
		return true
	}
}

func (l *Link) Connect() {
	logPrefix := "[Connect]"
	res := l.FSM.Event(EventConnect)
	if res == false {
		log.Printf("%v Fail to EventConnect", logPrefix)
	} else {
		log.Printf("%v Current[%v]", logPrefix, getStateString(l.FSM.Current()))
	}
}

func (l *Link) Echo() {
	logPrefix := "[Echo]"
	res := l.FSM.Event(EventEcho)
	if res == false {
		log.Printf("%v Fail to EventEcho", logPrefix)
	} else {
		log.Printf("%v Current[%v]", logPrefix, getStateString(l.FSM.Current()))
	}
}

func (l *Link) Disconnect() {
	logPrefix := "[Disconnect]"
	res := l.FSM.Event(EventDisconnect)
	if res == false {
		log.Printf("%v Fail to EventDisconnect", logPrefix)
	} else {
		log.Printf("%v Current[%v]", logPrefix, getStateString(l.FSM.Current()))
	}
}

func main() {
	link := NewLink(StateDisconnected)
	link.Connect()
	link.Echo()
	link.Echo()
	link.Echo()
	link.Disconnect()
}
