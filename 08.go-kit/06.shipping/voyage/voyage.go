package voyage

import (
	"errors"
	"time"

	"08.go-kit/06.shipping/location"
)

type Number string

type Voyage struct {
	Number   Number
	Schedule Schedule
}

func New(n Number, s Schedule) *Voyage {
	return &Voyage{Number: n, Schedule: s}
}

type Schedule struct {
	CarrierMovements []CarrierMovement
}

type CarrierMovement struct {
	DepartureLocation location.UNLocode
	ArrivalLocation   location.UNLocode
	DepartureTime     time.Time
	ArrivalTime       time.Time
}

var ErrUnknown = errors.New("unknown voyage")

type Repository interface {
	Find(Number) (*Voyage, error)
}
