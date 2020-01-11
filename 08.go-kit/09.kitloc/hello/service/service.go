package service

import (
	"context"
	"errors"
)

type Service struct {
}

func (s Service) Hello(ctx context.Context, firstName string, lastName string) (string, error) {
	panic(errors.New("not implemented"))
}
