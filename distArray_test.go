package DistributedArray_test

import (
	"fmt"
	"testing"

	da "github.com/RossMerr/Caudex.DistributedArray"
)

func Test_NewDistArray(t *testing.T) {

	pg := da.NewPlaceGroup()

	distArray := da.NewDistArray(da.NewDistributionStratege(da.Unique, pg))

	fmt.Printf("%+v", distArray)
}
