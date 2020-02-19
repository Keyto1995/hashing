package interfaces

import "github.com/keyto1995/hashing/hashcode"

type HashFunction interface {
	NewHasher() Hasher

	HashString(input string) hashcode.HashCode
	HashBytes(input []byte) hashcode.HashCode
}
