package mm32

import (
	"bytes"
	"github.com/keyto1995/hashing/hashcode"
	"github.com/keyto1995/hashing/interfaces"
	"math"
)

type Hasher struct {
	h1     uint32
	shift  int
	length int
	buffer uint64
}

func NewHasher(seed uint32) *Hasher {
	return &Hasher{
		h1:     seed,
		length: 0,
	}
}

func (h *Hasher) update(nBytes int, update uint64) {
	// 1 <= nBytes <= 4
	h.buffer |= update << h.shift
	h.shift += nBytes * 8
	h.length += nBytes
	if h.shift >= 32 {
		h.h1 = mixH1(h.h1, mixK1(uint32(h.buffer)))
		h.buffer >>= 32
		h.shift -= 32
	}
}

func (h *Hasher) PutByte(b byte) interfaces.Hasher {
	h.update(1, uint64(b))
	return h
}

func (h *Hasher) PutBool(b bool) interfaces.Hasher {
	if b {
		return h.PutByte(1)
	} else {
		return h.PutByte(0)
	}
}

func (h *Hasher) PutBytes(bs []byte) interfaces.Hasher {
	for _, v := range bs {
		h.PutByte(v)
	}
	return h
}

func (h *Hasher) PutUint32(i uint32) interfaces.Hasher {
	h.update(4, uint64(i))
	return h
}

func (h *Hasher) PutUint64(i uint64) interfaces.Hasher {
	h.PutUint32(uint32(i))
	h.PutUint32(uint32(i >> 32))
	return h
}

func (h *Hasher) PutFloat32(f float32) interfaces.Hasher {
	h.PutUint32(math.Float32bits(f))
	return h
}

func (h *Hasher) PutFloat64(f float64) interfaces.Hasher {
	return h.PutUint64(math.Float64bits(f))
}

func (h *Hasher) PutString(s string) interfaces.Hasher {
	b := bytes.NewBufferString(s)
	return h.PutBytes(b.Bytes())
}

func (h *Hasher) Hash() hashcode.HashCode {
	h.h1 ^= mixK1(uint32(h.buffer))
	return fmix(h.h1, h.length)
}
