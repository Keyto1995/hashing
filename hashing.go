package hashing

import (
	"github.com/keyto1995/hashing/interfaces"
	"github.com/keyto1995/hashing/murmur3/mm32"
)

func Murmur32() interfaces.HashFunction {
	return mm32.NewHashFunction(0)
}
