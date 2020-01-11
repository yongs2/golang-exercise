package kitloc

import "context"

type Service interface {
	Hello(ctx context.Context, firstName string, lastName string) (greeting string, err error)
}
