package DistributedArray

import "sync"

// Place representation from the APGAS model
type Place struct {
	id    int
	array *DistArray
}

func (s *Place) Id() int {
	return s.id
}

// PlaceGroup represents an ordered set of Places
type PlaceGroup struct {
	places []Place
}

// NewPlaceGroup returns a DistributedArray.PlaceGroup
func NewPlaceGroup() *PlaceGroup {
	return &PlaceGroup{}
}

func (s *PlaceGroup) Async(f func(Place)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(s.places))
	for _, p := range s.places {
		go func() {
			f(p)
			wg.Done()
		}()
	}

	wg.Wait()
}

// Iterator iterates through all places
func (s *PlaceGroup) Iterator() *PlaceIterator {
	i := &PlaceIterator{
		group: s,
		last:  0,
	}
	return i
}

type PlaceIterator struct {
	group *PlaceGroup
	last  int
}

// HasNext checks the iterator has any more values
func (s *PlaceIterator) HasNext() bool {
	if s.last >= len(s.group.places) {
		return false
	}
	return true
}

// Async is the denotational mechanism to express activities that perform computation in a place
func (s *PlaceIterator) Async(f func(Place)) {
	value := s.group.places[s.last]
	f(value)
	s.last++
}
