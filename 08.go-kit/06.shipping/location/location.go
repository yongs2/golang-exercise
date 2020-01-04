package location

import (
	"errors"
)

type UNLocode string

type Location struct {
	UNLocode UNLocode
	Name     string
}

var ErrUnknown = errors.New("unknown location")

type Repository interface {
	Find(locode UNLocode) (*Location, error)
	FindAll() []*Location
}
