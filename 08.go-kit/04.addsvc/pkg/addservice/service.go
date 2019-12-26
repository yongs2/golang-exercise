package addservice

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

// Service ...
type Service interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// New ...
func New(logger log.Logger, ints, chars metrics.Counter) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware(ints, chars)(svc)
	}
	return svc
}

var (
	// ErrTwoZeroes ...
	ErrTwoZeroes = errors.New("cann't sum two zeroes")

	// ErrIntOverflow ...
	ErrIntOverflow = errors.New("integer overflow")

	// ErrMaxSizeExceeded ...
	ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)

// NewBasicService ...
func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

func (s basicService) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

func (s basicService) Concat(_ context.Context, a, b string) (string, error) {
	if len(a)+len(b) > maxLen {
		return "", ErrMaxSizeExceeded
	}
	return a + b, nil
}
