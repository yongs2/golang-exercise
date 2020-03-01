package main

import (
	"fmt"
	"time"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}
	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state":  func(e *fsm.Event) { d.enterState(e) },
			"after_state":  func(e *fsm.Event) { d.afterState(e) },
			"leave_state":  func(e *fsm.Event) { d.leaveState(e) },
			"before_event": func(e *fsm.Event) { d.beforeEvent(e) },
			"after_event":  func(e *fsm.Event) { d.afterEvent(e) },
		},
	)
	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("%s >> enterState(%s), dst=%s\n", time.Now().Format("15:04:05.999"), d.To, e.Dst)
}

func (d *Door) afterState(e *fsm.Event) {
	fmt.Printf("%s >> afterState(%s), dst=%s\n", time.Now().Format("15:04:05.999"), d.To, e.Dst)
}

func (d *Door) leaveState(e *fsm.Event) {
	fmt.Printf("%s >> leaveState(%s), src=%s\n", time.Now().Format("15:04:05.999"), d.To, e.Src)
}

func (d *Door) beforeEvent(e *fsm.Event) {
	fmt.Printf("%s > beforeEvent(%s), src=%s\n", time.Now().Format("15:04:05.999"), d.To, e.Src)
}

func (d *Door) afterEvent(e *fsm.Event) {
	fmt.Printf("%s > afterEvent(%s), src=%s\n", time.Now().Format("15:04:05.999"), d.To, e.Src)
}

func main() {
	door := NewDoor("heaven")

	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
