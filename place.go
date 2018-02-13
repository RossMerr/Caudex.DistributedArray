package DistributedArray

// Place representation from the APGAS model
type Place struct {
	id    int64
	array *DistArray
}

func (s *Place) Id() int64 {
	return s.id
}

// PlaceGroup represents an ordered set of Places
type PlaceGroup struct {
	places []*Place
}

// NewPlaceGroup returns a DistributedArray.PlaceGroup
func NewPlaceGroup() *PlaceGroup {
	return &PlaceGroup{}
}

// Broadcast send a update to all places
func (s *PlaceGroup) Broadcast(f func(*Place)) {
	for _, p := range s.places {
		f(p)
	}
}
