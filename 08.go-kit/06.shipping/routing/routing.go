package routing

import (
	"08.go-kit/06.shipping/cargo"
)

type Service interface {
	FetchRoutesForSpecification(rs cargo.RouteSpecification) []cargo.Itinerary
}
