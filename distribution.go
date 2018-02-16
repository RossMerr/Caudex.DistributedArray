package DistributedArray

import (
	"fmt"
)

// Distribution stratege for initializing the distributed array
type Distribution interface {
	Stratege()
	Places() []Place
}

// DistributionStratege name of the distribution stratege
type DistributionStratege string

// NewDistributionStratege returns a distribution stratege for the distributed array
func NewDistributionStratege(name DistributionStratege, pg *PlaceGroup) Distribution {
	if name == Unique {
		return &UniqueDistribution{pg: pg}
	}

	panic(fmt.Sprintf("Distribution stratege not found %+v", name))
}

const (
	// Unique name of the unique distribution
	Unique DistributionStratege = "UniqueDistribution"
)

// UniqueDistribution initializes the distributed array with unique distribution
type UniqueDistribution struct {
	pg *PlaceGroup
}

// Stratege this is just a placeholder method
func (s *UniqueDistribution) Stratege() {

}

func (s *UniqueDistribution) Places() []Place {
	return s.pg.places
}
