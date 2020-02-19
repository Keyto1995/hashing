package interfaces

import "github.com/keyto1995/hashing/hashcode"

// Hasher
type Hasher interface {
	PutByte(b byte) Hasher
	PutBool(b bool) Hasher
	PutBytes(bytes []byte) Hasher
	PutUint32(i uint32) Hasher
	PutUint64(i uint64) Hasher
	PutFloat32(f float32) Hasher
	PutFloat64(f float64) Hasher
	PutString(s string) Hasher
	Hash() hashcode.HashCode
}
