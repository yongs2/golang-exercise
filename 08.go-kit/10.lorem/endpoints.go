package lorem

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type LoremRequest struct {
	RequestType string
	Min         int
	Max         int
}

type LoremResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

func MakeLoremEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		var (
			txt      string
			min, max int
		)

		min = req.Min
		max = req.Max

		if strings.EqualFold(req.RequestType, "Word") {
			txt = svc.Word(min, max)
		} else if strings.EqualFold(req.RequestType, "Sentence") {
			txt = svc.Sentence(min, max)
		} else if strings.EqualFold(req.RequestType, "Paragraph") {
			txt = svc.Paragraph(min, max)
		} else {
			return nil, ErrRequestTypeNotFound
		}

		return LoremResponse{Message: txt}, nil
	}

}
