package mm128

import (
	"encoding/binary"
	"github.com/keyto1995/hashing/hashcode"
	"github.com/keyto1995/hashing/interfaces"
	"math"
	"math/bits"
	"unsafe"
)

type Hasher struct {
	clen int    // Digested input cumulative length.
	tail []byte // 0 to Size()-1 bytes view of `buf'.
	buf  [chunkSize]byte
	h1   uint64 // Unfinalized running hash part 1.
	h2   uint64 // Unfinalized running hash part 2.
}

func NewHasher(seed uint32) interfaces.Hasher {
	h := &Hasher{
		h1: uint64(seed),
		h2: uint64(seed),
	}
	return h
}

func (h *Hasher) update(p []byte) {
	h.clen += len(p)

	if len(h.tail) > 0 {
		// Stick back pending bytes.
		nfree := h.Size() - len(h.tail) // nfree ∈ [1, h.Size()-1].
		if nfree < len(p) {
			// One full block can be formed.
			block := append(h.tail, p[:nfree]...)
			p = p[nfree:]
			_ = h.bmix(block) // No tail.
		} else {
			// Tail's buf is large enough to prevent reallocs.
			p = append(h.tail, p...)
		}
	}

	h.tail = h.bmix(p)

	// Keep own copy of the 0 to Size()-1 pending bytes.
	nn := copy(h.buf[:], h.tail)
	h.tail = h.buf[:nn]

}

func (h *Hasher) bmix(p []byte) (tail []byte) {
	h1, h2 := h.h1, h.h2

	nblocks := len(p) / h.Size()
	for i := 0; i < nblocks; i++ {
		t := (*[2]uint64)(unsafe.Pointer(&p[i*16]))
		k1, k2 := t[0], t[1]

		h1 ^= mixK1(k1)

		h1 = bits.RotateLeft64(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729

		h2 ^= mixK2(k2)

		h2 = bits.RotateLeft64(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}
	h.h1, h.h2 = h1, h2
	return p[nblocks*h.Size():]
}

func (h *Hasher) sum128() (h1, h2 uint64) {

	h1, h2 = h.h1, h.h2

	var k1, k2 uint64
	switch len(h.tail) % h.Size() {
	case 15:
		k2 ^= uint64(h.tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(h.tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(h.tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(h.tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(h.tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(h.tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(h.tail[8]) << 0

		h2 ^= mixK2(k2)
		fallthrough

	case 8:
		k1 ^= uint64(h.tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(h.tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(h.tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(h.tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(h.tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(h.tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(h.tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(h.tail[0]) << 0

		h1 ^= mixK1(k1)
	}

	h1 ^= uint64(h.clen)
	h2 ^= uint64(h.clen)

	h1 += h2
	h2 += h1

	h1 = fmix64(h1)
	h2 = fmix64(h2)

	h1 += h2
	h2 += h1

	return h1, h2
}

func (h *Hasher) PutByte(b byte) interfaces.Hasher {
	h.update([]byte{b})
	return h
}

func (h *Hasher) PutBool(b bool) interfaces.Hasher {
	var bi byte
	if b {
		bi = 1
	} else {
		bi = 0
	}
	h.update([]byte{bi})
	return h
}

func (h *Hasher) PutBytes(bytes []byte) interfaces.Hasher {
	h.update(bytes)
	return h
}

func (h *Hasher) PutUint32(i uint32) interfaces.Hasher {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	h.update(b)
	return h
}

func (h *Hasher) PutUint64(i uint64) interfaces.Hasher {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	h.update(b)
	return h
}

func (h *Hasher) PutFloat32(f float32) interfaces.Hasher {
	return h.PutUint32(math.Float32bits(f))
}

func (h *Hasher) PutFloat64(f float64) interfaces.Hasher {
	return h.PutUint64(math.Float64bits(f))
}

func (h *Hasher) PutString(s string) interfaces.Hasher {
	b := []byte(s)
	return h.PutBytes(b)
}

func (h *Hasher) Hash() hashcode.HashCode {
	h1, h2 := h.sum128()
	var b = make([]byte, chunkSize)
	binary.LittleEndian.PutUint64(b[0:8], h1)
	binary.LittleEndian.PutUint64(b[8:16], h2)
	return hashcode.NewBytesHashCode(b)
}

func (h *Hasher) Size() int {
	return chunkSize
}
