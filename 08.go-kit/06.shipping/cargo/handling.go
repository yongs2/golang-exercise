package cargo

import (
	"errors"
	"time"

	"08.go-kit/06.shipping/location"
	"08.go-kit/06.shipping/voyage"
)

type HandlingActivity struct {
	Type         HandlingEventType
	Location     location.UNLocode
	VoyageNumber voyage.Number
}

type HandlingEvent struct {
	TrackingID TrackingID
	Activity   HandlingActivity
}

type HandlingEventType int

const (
	NotHandled HandlingEventType = iota
	Load
	Unload
	Receive
	Claim
	Customs
)

func (t HandlingEventType) String() string {
	switch t {
	case NotHandled:
		return "Not Handled"
	case Load:
		return "Load"
	case Unload:
		return "Unload"
	case Receive:
		return "Receive"
	case Claim:
		return "Claim"
	case Customs:
		return "Customs"
	}
	return ""
}

type HandlingHistory struct {
	HandlingEvents []HandlingEvent
}

func (h HandlingHistory) MostRecentlyCompletedEvent() (HandlingEvent, error) {
	if len(h.HandlingEvents) == 0 {
		return HandlingEvent{}, errors.New("delivery history is empty")
	}
	return h.HandlingEvents[len(h.HandlingEvents)-1], nil
}

type HandlingEventRepository interface {
	Store(e HandlingEvent)
	QueryHandlingHistory(TrackingID) HandlingHistory
}

type HandlingEventFactory struct {
	CargoRepository    Repository
	VoyageRepository   voyage.Repository
	LocationRepository location.Repository
}

func (f *HandlingEventFactory) CreateHandlingEvent(registered time.Time, completed time.Time, id TrackingID,
	voyageNumber voyage.Number, unLocode location.UNLocode, eventType HandlingEventType) (HandlingEvent, error) {

	if _, err := f.CargoRepository.Find(id); err != nil {
		return HandlingEvent{}, err
	}

	if _, err := f.VoyageRepository.Find(voyageNumber); err != nil {
		// TODO: This is pretty ugly, but when creating a Receive event, the voyage number is not known.
		if len(voyageNumber) > 0 {
			return HandlingEvent{}, err
		}
	}

	if _, err := f.LocationRepository.Find(unLocode); err != nil {
		return HandlingEvent{}, err
	}

	return HandlingEvent{
		TrackingID: id,
		Activity: HandlingActivity{
			Type:         eventType,
			Location:     unLocode,
			VoyageNumber: voyageNumber,
		},
	}, nil
}
