package handling

import (
	"errors"
	"time"

	"08.go-kit/06.shipping/cargo"
	"08.go-kit/06.shipping/inspection"
	"08.go-kit/06.shipping/location"
	"08.go-kit/06.shipping/voyage"
)

var ErrInvalidArgument = errors.New("invalid argument")

type EventHandler interface {
	CargoWasHandled(cargo.HandlingEvent)
}

type Service interface {
	RegisterHandlingEvent(completed time.Time, id cargo.TrackingID, voyageNumber voyage.Number, unLocode location.UNLocode, eventType cargo.HandlingEventType) error
}

type service struct {
	handlingEventRepository cargo.HandlingEventRepository
	handlingEventFactory    cargo.HandlingEventFactory
	handlingEventHandler    EventHandler
}

func (s *service) RegisterHandlingEvent(completed time.Time, id cargo.TrackingID, voyageNumber voyage.Number, loc location.UNLocode, eventType cargo.HandlingEventType) error {
	if completed.IsZero() || id == "" || loc == "" || eventType == cargo.NotHandled {
		return ErrInvalidArgument
	}

	e, err := s.handlingEventFactory.CreateHandlingEvent(time.Now(), completed, id, voyageNumber, loc, eventType)
	if err != nil {
		return err
	}

	s.handlingEventRepository.Store(e)
	s.handlingEventHandler.CargoWasHandled(e)
	return nil
}

func NewService(r cargo.HandlingEventRepository, f cargo.HandlingEventFactory, h EventHandler) Service {
	return &service{
		handlingEventRepository: r,
		handlingEventFactory:    f,
		handlingEventHandler:    h,
	}
}

type handlingEventHandler struct {
	InspectionService inspection.Service
}

func (h *handlingEventHandler) CargoWasHandled(event cargo.HandlingEvent) {
	h.InspectionService.InspectCargo(event.TrackingID)
}

func NewEventHandler(s inspection.Service) EventHandler {
	return &handlingEventHandler{
		InspectionService: s,
	}
}
