package DistributedArray

import (
	"fmt"
)

// Place representation from the APGAS model
type Place struct {
	id int64
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

//
func (s *PlaceGroup) Broadcast() {
	for _, p := range s.places {
		fmt.Print(p.Id())
	}
}
