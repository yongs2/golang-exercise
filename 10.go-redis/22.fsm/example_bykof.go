package main

import (
	"errors"
	"log"

	"github.com/bykof/stateful"
)

type (
	MyMachine struct {
		state  stateful.State
		amount int
	}

	AmountParams struct {
		Amount int
	}
)

func NewMyMachine() MyMachine {
	return MyMachine{
		state:  A,
		amount: 0,
	}
}

func (mm MyMachine) State() stateful.State {
	return mm.state
}

func (mm *MyMachine) SetState(state stateful.State) error {
	mm.state = state
	return nil
}

const (
	A = stateful.DefaultState("A")
	B = stateful.DefaultState("B")
)

func (mm *MyMachine) FromAToB(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	amountParams, ok := transitionArguments.(AmountParams)
	if !ok {
		return nil, errors.New("could not parse AmountParams")
	}

	mm.amount += amountParams.Amount
	return B, nil
}

func (mm *MyMachine) FromBToA(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	amountParams, ok := transitionArguments.(AmountParams)
	if !ok {
		return nil, errors.New("could not parse AmountParams")
	}

	mm.amount -= amountParams.Amount
	return A, nil
}

func (mm *MyMachine) FromAToNotExistingC(_ stateful.TransitionArguments) (stateful.State, error) {
	return stateful.DefaultState("C"), nil
}

func main() {
	myMachine := NewMyMachine()
	stateMachine := &stateful.StateMachine{
		StatefulObject: &myMachine,
	}

	stateMachine.AddTransition(
		// The transition function
		myMachine.FromAToB,
		// SourceStates
		stateful.States{A},
		// DestinationStates
		stateful.States{B},
	)

	stateMachine.AddTransition(
		myMachine.FromBToA,
		stateful.States{B},
		stateful.States{A},
	)

	log.Printf("myMachine[state=%v, amount=%d]\n", myMachine.state, myMachine.amount)

	var err error
	err = stateMachine.Run(
		// The transition function
		myMachine.FromAToB,
		// The transition params which will be passed to the transition function
		stateful.TransitionArguments(AmountParams{Amount: 1}),
	)
	log.Printf("myMachine[state=%v, amount=%d], FromAtoB=%v\n", myMachine.state, myMachine.amount, err)

	err = stateMachine.Run(
		myMachine.FromBToA,
		stateful.TransitionArguments(AmountParams{Amount: 1}),
	)
	log.Printf("myMachine[state=%v, amount=%d], FromBToA=%v\n", myMachine.state, myMachine.amount, err)

	err = stateMachine.Run(
		myMachine.FromBToA,
		stateful.TransitionArguments(AmountParams{Amount: 1}),
	)
	log.Printf("myMachine[state=%v, amount=%d], FromBToA=%v\n", myMachine.state, myMachine.amount, err)

	err = stateMachine.Run(
		myMachine.FromAToNotExistingC,
		stateful.TransitionArguments(nil),
	)
	log.Printf("myMachine[state=%v, amount=%d], FromAToNotExistingC=%v\n", myMachine.state, myMachine.amount, err)
}
