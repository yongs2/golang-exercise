package inspection

import (
	"08.go-kit/06.shipping/cargo"
)

type EventHandler interface {
	CargoWasMisdirected(*cargo.Cargo)
	CargoHasArrived(*cargo.Cargo)
}

type Service interface {
	InspectCargo(id cargo.TrackingID)
}

type service struct {
	cargos  cargo.Repository
	events  cargo.HandlingEventRepository
	handler EventHandler
}

func (s *service) InspectCargo(id cargo.TrackingID) {
	c, err := s.cargos.Find(id)
	if err != nil {
		return
	}

	h := s.events.QueryHandlingHistory(id)

	c.DeriveDeliveryProgress(h)

	if c.Delivery.IsMisdirected {
		s.handler.CargoWasMisdirected(c)
	}

	if c.Delivery.IsUnloadedAtDestination {
		s.handler.CargoHasArrived(c)
	}

	s.cargos.Store(c)
}

func NewService(cargos cargo.Repository, events cargo.HandlingEventRepository, handler EventHandler) Service {
	return &service{cargos, events, handler}
}
