package DistributedArray

// DistArray a 1 dimensional distributed array
type DistArray struct {
	pg            *PlaceGroup
	length        int
	minLocalIndex int
	maxLocalIndex int
	internalArray []interface{}
}

// NewDistArray returns a DistributedArray.DistArray
func NewDistArray(length int, pg *PlaceGroup) *DistArray {
	return &DistArray{pg: pg, length: length, internalArray: make([]interface{}, length)}
}

// Set the given index of the array with the element
func (s *DistArray) Set(i int, v interface{}) {
	if i < s.minLocalIndex || i > s.maxLocalIndex {
		s.internalArray[i] = v
	}
}

// At returns the element of this array from the given index
func (s *DistArray) At(i int) interface{} {
	if i < s.minLocalIndex || i > s.maxLocalIndex {
		return s.internalArray[i]
	}

	return nil
}

// Reduce applies a function against each element in the array to reduce it to a single value
func (s *DistArray) Reduce(f ReducerFunc) {
}

// Map method creates a new array from calling a function on every element in the array
func (s *DistArray) Map(f MapperFunc) {
	s.pg.Broadcast(func(p *Place) {
		p.array.Map(f)
	})
}

// LocalIterator iterates through all local elements
func (s *DistArray) LocalIterator() Iterator {
	i := &DistArrayIterator{
		array: s,
		last:  0,
	}
	return i
}

// GlobalIterator iterates through all global elements
func (s *DistArray) GlobalIterator() Iterator {
	i := &DistArrayIterator{
		array: s,
		last:  0,
	}
	return i
}

type DistArrayIterator struct {
	array *DistArray
	last  int
}

// HasNext checks the iterator has any more values
func (s *DistArrayIterator) HasNext() bool {
	if s.last >= s.array.length {
		return false
	}
	return true
}

// Next moves the iterator and returns the next element
func (s *DistArrayIterator) Next() interface{} {
	value := s.array.internalArray[s.last]
	s.last++
	return value
}
