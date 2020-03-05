package main

import (
	"log"
	"time"

	"github.com/ryanfaerman/fsm"
)

type Thing struct {
	State fsm.State

	machine *fsm.Machine
}

func (t *Thing) CurrentState() fsm.State { return t.State }
func (t *Thing) SetState(s fsm.State) { t.State = s }
func (t *Thing) Apply(r *fsm.Ruleset) *fsm.Machine {
    if t.machine == nil {
        t.machine = &fsm.Machine{Subject: t}
    }

    t.machine.Rules = r
    return t.machine
}

func main() {
	var err error

	rules := fsm.Ruleset{}
	rules.AddTransition(fsm.T{"pending", "started"})
	rules.AddTransition(fsm.T{"started", "finished"})
	rules.AddTransition(fsm.T{"finished", "pending"})
	rules.AddRule(fsm.T{"started", "finished"}, func(subject fsm.Stater, goal fsm.State) bool {
		time.Sleep(1 * time.Second)
		log.Println("sleep 1")
		return true
	})

	var some_thing Thing
	events := []fsm.State{"pending", "started", "finished", "pending", "started", "finished"}
	for index, event := range events {
		if index == 0 {
			some_thing = Thing{State: "pending"}
		} else {
			err = some_thing.Apply(&rules).Transition(event)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Println(index, ":", event, ": ", rules.Permitted(&some_thing, "finished"))
		log.Println(index, ":", event, ": ", some_thing)
	}
}
