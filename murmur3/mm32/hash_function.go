package mm32

import (
	"bytes"
	"encoding/binary"
	"github.com/keyto1995/hashing/hashcode"
	"github.com/keyto1995/hashing/interfaces"
	"math/bits"
)

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
)

func NewHashFunction(seed uint32) *HashFunction {
	return &HashFunction{seed: seed}
}

type HashFunction struct {
	seed uint32
}

func (m *HashFunction) NewHasher() interfaces.Hasher {
	return NewHasher(0)
}

func (m *HashFunction) HashString(input string) hashcode.HashCode {
	bs := bytes.NewBufferString(input).Bytes()
	return m.HashBytes(bs)
}

func (m *HashFunction) HashBytes(input []byte) hashcode.HashCode {
	h1 := m.seed
	var i int
	var k1 uint32
	length := len(input)
	for i = 0; i+4 <= length; i += 4 {
		k1 = mixK1(getIntLittleEndian(input[i : i+4]))
		h1 = mixH1(h1, k1)
	}

	k1 = 0
	for shift := 0; i < length; shift += 8 {
		k1 ^= uint32(input[i]) << shift
		i++
	}

	h1 ^= mixK1(k1)
	hash := fmix(h1, length)
	return hashcode.NewUint32HashCode(hash)
}

func getIntLittleEndian(input []byte) uint32 {
	return binary.LittleEndian.Uint32(input)
}

func mixK1(k1 uint32) uint32 {
	k1 *= c1
	k1 = bits.RotateLeft32(k1, 15)
	k1 *= c2
	return k1
}

func mixH1(h1, k1 uint32) uint32 {
	h1 ^= k1
	h1 = bits.RotateLeft32(h1, 13)
	h1 = h1*5 + 0xe6546b64
	return h1
}

func fmix(h1 uint32, length int) uint32 {
	h1 ^= uint32(length)
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16
	return h1
}
