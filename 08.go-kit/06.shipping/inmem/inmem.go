package inmem

import (
	"sync"

	"08.go-kit/06.shipping/cargo"
	"08.go-kit/06.shipping/location"
	"08.go-kit/06.shipping/voyage"
)

type cargoRepository struct {
	mtx    sync.RWMutex
	cargos map[cargo.TrackingID]*cargo.Cargo
}

func (r *cargoRepository) Store(c *cargo.Cargo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.cargos[c.TrackingID] = c
	return nil
}

func (r *cargoRepository) Find(id cargo.TrackingID) (*cargo.Cargo, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.cargos[id]; ok {
		return val, nil
	}
	return nil, cargo.ErrUnknown
}

func (r *cargoRepository) FindAll() []*cargo.Cargo {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*cargo.Cargo, 0, len(r.cargos))
	for _, val := range r.cargos {
		c = append(c, val)
	}
	return c
}

func NewCargoRepository() cargo.Repository {
	return &cargoRepository{
		cargos: make(map[cargo.TrackingID]*cargo.Cargo),
	}
}

type locationRepository struct {
	locations map[location.UNLocode]*location.Location
}

func (r *locationRepository) Find(locode location.UNLocode) (*location.Location, error) {
	if l, ok := r.locations[locode]; ok {
		return l, nil
	}
	return nil, location.ErrUnknown
}

func (r *locationRepository) FindAll() []*location.Location {
	l := make([]*location.Location, 0, len(r.locations))
	for _, val := range r.locations {
		l = append(l, val)
	}
	return l
}

func NewLocationRepository() location.Repository {
	r := &locationRepository{
		locations: make(map[location.UNLocode]*location.Location),
	}

	r.locations[location.SESTO] = location.Stockholm
	r.locations[location.AUMEL] = location.Melbourne
	r.locations[location.CNHKG] = location.Hongkong
	r.locations[location.JNTKO] = location.Tokyo
	r.locations[location.NLRTM] = location.Rotterdam
	r.locations[location.DEHAM] = location.Hamburg

	return r
}

type voyageRepository struct {
	voyages map[voyage.Number]*voyage.Voyage
}

func (r *voyageRepository) Find(voyageNumber voyage.Number) (*voyage.Voyage, error) {
	if v, ok := r.voyages[voyageNumber]; ok {
		return v, nil
	}
	return nil, voyage.ErrUnknown
}

func NewVoyageRepository() voyage.Repository {
	r := &voyageRepository{
		voyages: make(map[voyage.Numer]*voyage.Voyage),
	}

	r.voyages[voyage.V100.Number] = voyage.V100
	r.voyages[voyage.V300.Number] = voyage.V300
	r.voyages[voyage.V400.Number] = voyage.V400

	r.voyages[voyage.V0100S.Number] = voyage.V0100S
	r.voyages[voyage.V0200T.Number] = voyage.V0200T
	r.voyages[voyage.V0300A.Number] = voyage.V0300A
	r.voyages[voyage.V0301S.Number] = voyage.V0301S
	r.voyages[voyage.V0400S.Number] = voyage.V0400S

	return r
}

type handlingEventRepository struct {
	mtx    sync.RWMutex
	events map[cargo.TrackingID][]cargo.HandlingEvent
}

func (r *handlingEventRepository) Store(e cargo.HandlingEvent) {
	r.mtx.Lock()
	defer r.mtx.Unlokc()

	if _, ok := r.events[e.TrackingID]; !ok {
		r.events[e.TrackingID] = make([]cargo.HandlingEvent, 0)
	}
	r.events[e.TrackingID] = append(r.event[e.TrackingID], e)
}

func (r *handlingEventRepository) QueryHandlingHistory(id cargo.TrackingID) cargo.HandlingHistory {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return cargo.HandlingHistory{HandlingEvents: r.events[id]}
}

func NewHandlingEventRepository() cargo.HandlingEventRepository {
	return &handlingEventRepository{
		events: make(map[cargo.TrackingID][]cargo.HandlingEvent),
	}
}
