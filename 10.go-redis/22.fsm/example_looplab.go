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
	events := []string{"open", "close", "open", "close"}
	for _, event := range events {
		err := door.FSM.Event(event)
		if err != nil {
			fmt.Println(err)
		}
	}
	
	fsm1 := fsm.NewFSM(
		"E_FS_IDLE",
		fsm.Events{
			// MT
			{Name: "EVT_MT_REQ", Src: []string{"E_FS_IDLE"}, Dst: "E_FS_OPEN"},
			{Name: "No-EBC-DT", Src: []string{"E_FS_OPEN"}, Dst: "E_FS_DT_OPEN"},
			{Name: "with-EBC-MT", Src: []string{"E_FS_OPEN"}, Dst: "E_FS_MT_OPEN"},
			{Name: "EVT_DAA-SUCC", Src: []string{"E_FS_DT_OPEN"}, Dst: "E_FS_DT_WAIT_REPORT"},
			{Name: "EVT_DAA-TIMEOUT", Src: []string{"E_FS_DT_OPEN"}, Dst: "E_FS_DT_RETRY"},
			{Name: "EVT_DT-RETRY", Src: []string{"E_FS_DT_RETRY"}, Dst: "E_FS_DT_OPEN"},
			{Name: "EVT_DNR-SUCC", Src: []string{"E_FS_DT_WAIT_REPORT"}, Dst: "E_FS_DT_WAIT_ECB"},
			{Name: "EVT_DNR-FAIL", Src: []string{"E_FS_DT_WAIT_REPORT"}, Dst: "E_FS_DT_RETRY"},
			{Name: "EVT_CDR_ECB", Src: []string{"E_FS_DT_WAIT_ECB"}, Dst: "E_FS_DT_RETRY"},
			{Name: "No-EBC-Type-B", Src: []string{"E_FS_DT_WAIT_ECB"}, Dst: "E_FS_DT_SUCCESS"},
			{Name: "No-EBC-Type-O", Src: []string{"E_FS_DT_WAIT_ECB"}, Dst: "E_FS_DT_FAILURE"},
			{Name: "No-EBC-DT", Src: []string{"E_FS_MT_OPEN"}, Dst: "E_FS_DT_OPEN"},
			{Name: "EVT_TDA-SUCC", Src: []string{"E_FS_MT_OPEN"}, Dst: "E_FS_MT_SUCCESS"},
			{Name: "EVT_TDA-TIMEOUT", Src: []string{"E_FS_MT_OPEN"}, Dst: "E_FS_MT_RETRY"},
			{Name: "EVT_MT-RETRY", Src: []string{"E_FS_MT_RETRY"}, Dst: "E_FS_MT_OPEN"},
			{Name: "Check-Fail", Src: []string{"E_FS_DT_OPEN", "E_FS_DT_RETRY"}, Dst: "E_FS_DT_FAILURE"},
			{Name: "Check-Fail", Src: []string{"E_FS_MT_OPEN", "E_FS_MT_RETRY"}, Dst: "E_FS_MT_FAILURE"},
			{Name: "Close", Src: []string{"E_FS_DT_SUCCESS", "E_FS_DT_FAILURE", "E_FS_MT_SUCCESS", "E_FS_MT_FAILURE"}, Dst: "E_FS_CLOSE"},

			// MO
			{Name: "EVT_ODR", Src: []string{"E_FS_IDLE"}, Dst: "E_FS_OPEN"},
			{Name: "MO-Open", Src: []string{"E_FS_OPEN"}, Dst: "E_FS_MO_OPEN"},			
			{Name: "EVT_MO-NOTI", Src: []string{"E_FS_MO_OPEN"}, Dst: "E_FS_MO_SUCCESS"},
			{Name: "EVT_MO-NOTI-TIMEOUT", Src: []string{"E_FS_MO_OPEN"}, Dst: "E_FS_MO_RETRY"},			
			{Name: "EVT_MO-RETRY", Src: []string{"E_FS_MO_RETRY"}, Dst: "E_FS_MO_OPEN"},
			{Name: "Check-Fail", Src: []string{"E_FS_OPEN", "E_FS_MO_OPEN", "E_FS_MO_RETRY"}, Dst: "E_FS_MO_FAILURE"},
			{Name: "Close", Src: []string{"E_FS_MO_SUCCESS", "E_FS_MO_FAILURE"}, Dst: "E_FS_CLOSE"},
			{Name: "Idle", Src: []string{"E_FS_CLOSE"}, Dst: "E_FS_IDLE"},
		},
		fsm.Callbacks{
			"leave_state":  func(e *fsm.Event) {
				fmt.Printf("%s >> leaveState.src=%s, dst=%s\n", time.Now().Format("15:04:05.999"), e.Src, e.Dst)
			},
			"enter_state":  func(e *fsm.Event) {
				fmt.Printf("%s >> enterState.src=%s\n", time.Now().Format("15:04:05.999"), e.Src)
			},
		},
	)

	// MO
	fmt.Println(">> MO ------------")
	events = []string{
		"EVT_ODR", "MO-Open", "EVT_MO-NOTI", "Close", "Idle",
		"EVT_ODR", "Check-Fail", "Close", "Idle",
		"EVT_ODR", "MO-Open", "Check-Fail", "Close", "Idle",
		"EVT_ODR", "MO-Open", "EVT_MO-NOTI-TIMEOUT", "Check-Fail", "Close", "Idle",
		"EVT_ODR", "MO-Open", "EVT_MO-NOTI-TIMEOUT", "EVT_MO-RETRY", "Check-Fail", "Close", "Idle",
	}
	for _, event := range events {
		err := fsm1.Event(event)
		if err != nil {
			fmt.Println(err)
		}
	}

	// MT
	fmt.Println(">> MT ------------")
	events = []string{
		"EVT_MT_REQ", "with-EBC-MT", "EVT_TDA-SUCC", "Close", "Idle",
		"EVT_MT_REQ", "with-EBC-MT", "Check-Fail", "Close", "Idle",
		"EVT_MT_REQ", "with-EBC-MT", "EVT_TDA-TIMEOUT", "Check-Fail", "Close", "Idle",
		"EVT_MT_REQ", "with-EBC-MT", "EVT_TDA-TIMEOUT", "EVT_MT-RETRY", "Check-Fail", "Close", "Idle",
		"EVT_MT_REQ", "with-EBC-MT", "No-EBC-DT", "EVT_DAA-SUCC", "EVT_DNR-SUCC", "No-EBC-Type-B", "Close", "Idle",
	}
	for _, event := range events {
		err := fsm1.Event(event)
		if err != nil {
			fmt.Println(err)
		}
	}
}
