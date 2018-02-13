package DistributedArray

// DistArray a 1 dimensional distributed array
type DistArray struct {
	pg *PlaceGroup
}

// NewDistArray returns a DistributedArray.DistArray
func NewDistArray(pg *PlaceGroup) *DistArray {
	return &DistArray{pg: pg}
}
