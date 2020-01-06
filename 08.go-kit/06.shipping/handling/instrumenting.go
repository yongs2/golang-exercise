package handling

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"08.go-kit/06.shipping/cargo"
	"08.go-kit/06.shipping/location"
	"08.go-kit/06.shipping/voyage"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) RegisterHandlingEvent(completed time.Time, id cargo.TrackingID, voyageNumber voyage.Number,
	loc location.UNLocode, eventType cargo.HandlingEventType) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "register_incident").Add(1)
		s.requestLatency.With("method", "register_incident").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.RegisterHandlingEvent(completed, id, voyageNumber, loc, eventType)
}
