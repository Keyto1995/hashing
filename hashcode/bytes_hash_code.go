package hashcode

import (
	"encoding/binary"
	"encoding/hex"
)

type BytesHashCode struct {
	bytes []byte
}

func NewBytesHashCode(bytes []byte) *BytesHashCode {
	return &BytesHashCode{bytes: bytes}
}

func (b *BytesHashCode) Bits() int {
	return len(b.bytes) * 8
}

func (b *BytesHashCode) AsUint32() uint32 {
	return binary.LittleEndian.Uint32(b.bytes)
}

func (b *BytesHashCode) AsBytes() []byte {
	return b.bytes
}

func (b *BytesHashCode) AsHex() string {
	return hex.EncodeToString(b.bytes)
}
