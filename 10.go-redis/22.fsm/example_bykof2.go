package main

import (
	"errors"
	"log"

	"github.com/bykof/stateful"
	"github.com/bykof/stateful/statefulGraph"
)

const (
	E_FS_IDLE       = stateful.DefaultState("E_FS_IDLE")
	E_FS_OPEN       = stateful.DefaultState("E_FS_OPEN")
	E_FS_CLOSE      = stateful.DefaultState("E_FS_CLOSE")
	E_FS_MO_OPEN    = stateful.DefaultState("E_FS_MO_OPEN")
	E_FS_MO_SUCCESS = stateful.DefaultState("E_FS_MO_SUCCESS")
	E_FS_MO_FAILURE = stateful.DefaultState("E_FS_MO_FAILURE")
	E_FS_MO_RETRY   = stateful.DefaultState("E_FS_MO_RETRY")
)

type (
	MyMachine struct {
		state   stateful.State
		moState int
	}

	MoParams struct {
		MoState int
	}
)

func (mp *MoParams) GetData() interface{} {
	return mp
}

func NewMyMachine() MyMachine {
	return MyMachine{
		state:   E_FS_IDLE,
		moState: 0,
	}
}

func (mm MyMachine) State() stateful.State {
	return mm.state
}

func (mm *MyMachine) SetState(state stateful.State) error {
	mm.state = state
	return nil
}

func (mm *MyMachine) OnODR(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnODR(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_OPEN, nil
}

func (mm *MyMachine) OnMoOpen(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnMoOpen(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_MO_OPEN, nil
}

func (mm *MyMachine) OnMoNoti(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnMoNoti(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_MO_SUCCESS, nil
}

func (mm *MyMachine) OnMoNotiTimeout(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnMoNotiTimeout(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_MO_RETRY, nil
}

func (mm *MyMachine) OnMoRetry(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnMoRetry(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_MO_OPEN, nil
}

func (mm *MyMachine) OnCheckFail(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnCheckFail(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_MO_FAILURE, nil
}

func (mm *MyMachine) OnClose(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnClose(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_CLOSE, nil
}

func (mm *MyMachine) OnIdle(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	moParams, ok := transitionArguments.(MoParams)
	if !ok {
		return nil, errors.New("cannot parse moParams")
	}
	mm.moState = moParams.MoState
	log.Printf("[%v] OnIdle(MoState=%d).moState=%d\n", mm.state, moParams.MoState, mm.moState)
	return E_FS_IDLE, nil
}

func main() {
	myMachine := NewMyMachine()
	stateMachine := &stateful.StateMachine{
		StatefulObject: &myMachine,
	}

	// MO
	stateMachine.AddTransition(myMachine.OnODR, stateful.States{E_FS_IDLE}, stateful.States{E_FS_OPEN})
	stateMachine.AddTransition(myMachine.OnMoOpen, stateful.States{E_FS_OPEN}, stateful.States{E_FS_MO_OPEN})
	stateMachine.AddTransition(myMachine.OnMoNoti, stateful.States{E_FS_MO_OPEN}, stateful.States{E_FS_MO_SUCCESS})
	stateMachine.AddTransition(myMachine.OnMoNotiTimeout, stateful.States{E_FS_MO_OPEN}, stateful.States{E_FS_MO_RETRY})
	stateMachine.AddTransition(myMachine.OnMoRetry, stateful.States{E_FS_MO_RETRY}, stateful.States{E_FS_MO_OPEN})
	stateMachine.AddTransition(myMachine.OnCheckFail, stateful.States{E_FS_OPEN, E_FS_MO_OPEN, E_FS_MO_RETRY}, stateful.States{E_FS_MO_FAILURE})
	stateMachine.AddTransition(myMachine.OnClose, stateful.States{E_FS_MO_SUCCESS, E_FS_MO_FAILURE}, stateful.States{E_FS_CLOSE})
	stateMachine.AddTransition(myMachine.OnIdle, stateful.States{E_FS_CLOSE}, stateful.States{E_FS_IDLE})

	var err error

	// https://dreampuf.github.io/GraphvizOnline
	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: *stateMachine}
	err = stateMachineGraph.DrawGraph()
	log.Printf("stateMachineGraph.DrawGraph()=%v\n", err)

	// run the machine
	log.Printf("myMachine[state=%v, moState=%d]\n", myMachine.state, myMachine.moState)

	transitions := []stateful.Transition{
		myMachine.OnODR, myMachine.OnMoOpen, myMachine.OnMoNoti, myMachine.OnClose, myMachine.OnIdle,
		myMachine.OnODR, myMachine.OnCheckFail, myMachine.OnClose, myMachine.OnIdle,
		myMachine.OnODR, myMachine.OnMoOpen, myMachine.OnCheckFail, myMachine.OnClose, myMachine.OnIdle,
		myMachine.OnODR, myMachine.OnMoOpen, myMachine.OnMoNotiTimeout, myMachine.OnCheckFail, myMachine.OnClose, myMachine.OnIdle,
		myMachine.OnODR, myMachine.OnMoOpen, myMachine.OnMoNotiTimeout, myMachine.OnMoRetry, myMachine.OnCheckFail, myMachine.OnClose, myMachine.OnIdle,
	}
	for index, transition := range transitions {
		err = stateMachine.Run(transition, stateful.TransitionArguments(MoParams{MoState: index}))
		log.Printf("myMachine[state=%v, moState=%d], err=%v\n", myMachine.state, myMachine.moState, err)
	}
}
