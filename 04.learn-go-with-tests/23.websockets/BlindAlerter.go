package poker

import (
	"fmt"
	"time"
	"io"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	fmt.Println("BlindAlerterFunc.ScheduleAlertAt(", duration, ",", amount, ")")
	a(duration, amount, to)
}

func Alerter(duration time.Duration, amount int, to io.Writer) {
	fmt.Println("StdOutAlerter.duration=", duration, ", amount=", amount)
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
