package hashcode

import (
    "encoding/binary"
    "encoding/hex"
)

type Uint32HashCode struct {
    hash uint32
}

func NewUint32HashCode(hash uint32) HashCode {
    return &Uint32HashCode{hash: hash}
}

func (u *Uint32HashCode) Bits() int {
    return 32
}

func (u *Uint32HashCode) AsUint32() uint32 {
    return u.hash
}

func (u *Uint32HashCode) AsBytes() []byte {
    var bs = make([]byte, 4)
    binary.LittleEndian.PutUint32(bs, u.hash)
    return bs
}

func (u *Uint32HashCode) AsHex() string {
    return hex.EncodeToString(u.AsBytes())
}

