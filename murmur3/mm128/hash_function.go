package mm128

import (
	"encoding/binary"
	"github.com/keyto1995/hashing/hashcode"
	"github.com/keyto1995/hashing/interfaces"
	"math/bits"
	"unsafe"
)

const (
	chunkSize        = 16
	c1        uint64 = 0x87c37b91114253d5
	c2        uint64 = 0x4cf5ad432745937f
)

type HashFunction struct {
	seed uint32
}

func NewHashFunction(seed uint32) interfaces.HashFunction {
	return &HashFunction{seed: seed}
}

func (hf *HashFunction) NewHasher() interfaces.Hasher {
	return NewHasher(hf.seed)
}

func (hf *HashFunction) HashString(data string) hashcode.HashCode {
	bs := []byte(data)
	return hf.HashBytes(bs)
}

func (hf *HashFunction) HashBytes(data []byte) hashcode.HashCode {
	n := len(data)
	h1, h2 := uint64(hf.seed), uint64(hf.seed)

	var k1, k2 uint64
	nblocks := len(data) / chunkSize
	for i := 0; i < nblocks; i++ {
		t := (*[2]uint64)(unsafe.Pointer(&data[i*16]))
		k1, k2 = t[0], t[1]

		h1 ^= mixK1(k1)

		h1 = bits.RotateLeft64(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729

		h2 ^= mixK2(k2)

		h2 = bits.RotateLeft64(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}

	k1, k2 = 0, 0

	tail := data[nblocks*chunkSize:]

	switch len(tail) % chunkSize {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(tail[8]) << 0

		h2 ^= mixK2(k2)
		fallthrough

	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0]) << 0

		h1 ^= mixK1(k1)
	}

	h1 ^= uint64(n)
	h2 ^= uint64(n)

	h1 += h2
	h2 += h1

	h1 = fmix64(h1)
	h2 = fmix64(h2)

	h1 += h2
	h2 += h1

	var b = make([]byte, chunkSize)
	binary.LittleEndian.PutUint64(b[0:8], h1)
	binary.LittleEndian.PutUint64(b[8:16], h2)

	return hashcode.NewBytesHashCode(b)
}

func mixK1(k1 uint64) uint64 {
	k1 *= c1
	k1 = bits.RotateLeft64(k1, 31)
	k1 *= c2
	return k1
}

func mixK2(k2 uint64) uint64 {
	k2 *= c2
	k2 = bits.RotateLeft64(k2, 33)
	k2 *= c1
	return k2
}

func fmix64(k uint64) uint64 {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33
	return k
}
