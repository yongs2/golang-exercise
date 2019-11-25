package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	fmt.Println("BlindAlerterFunc.ScheduleAlertAt(", duration, ",", amount, ")")
	a(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	fmt.Println("StdOutAlerter.duration=", duration, ", amount=", amount)
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
