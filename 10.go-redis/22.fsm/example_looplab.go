package main

import (
	"fmt"
	"time"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(time.Now().Format("15:04:05.999"), fsm.Current())

	err := fsm.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Format("15:04:05.999"), fsm.Current())

	err = fsm.Event("close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Format("15:04:05.999"), fsm.Current())
}
