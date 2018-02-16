package DistributedArray_test

import (
	"fmt"
	"testing"

	da "github.com/RossMerr/Caudex.DistributedArray"
)

func Test_NewDistArray(t *testing.T) {
	pg := da.NewPlaceGroup()

	distArray := da.NewDistArray(da.NewDistributionStratege(da.Unique, pg))

	distArray.Set(1, 1)

	for iterator := distArray.LocalIterator(); iterator.HasNext(); {
		value := iterator.Next()

		fmt.Printf("%+v", value)

	}

	pg.Async(func(p da.Place) {
		fmt.Printf("%+v", p)
	})

	fmt.Printf("%+v", distArray)
}
