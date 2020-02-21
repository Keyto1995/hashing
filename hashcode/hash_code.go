package hashcode

type HashCode interface {
	// Bits HashCode编码长度
	Bits() int
	AsUint32() uint32
	AsBytes() []byte
	AsHex() string
}
