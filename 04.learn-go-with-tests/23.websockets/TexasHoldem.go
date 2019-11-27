package poker

import (
	"fmt"
	"time"
	"io/ioutil"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store: store,
	}
}

func (p *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	fmt.Println("TexasHoldem.Start(", numberOfPlayers, ")")
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, ioutil.Discard)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
