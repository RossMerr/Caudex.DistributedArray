package DistributedArray

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
