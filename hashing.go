package hashing

import (
	"github.com/keyto1995/hashing/interfaces"
	"github.com/keyto1995/hashing/murmur3/mm128"
	"github.com/keyto1995/hashing/murmur3/mm32"
)

// Murmur32 获得murmur3_32的hashfunction
func Murmur32() interfaces.HashFunction {
	return mm32.NewHashFunction(0)
}

// Murmur128 获得murmur3_128的hashfunction
func Murmur128() interfaces.HashFunction {
	return mm128.NewHashFunction(0)
}
