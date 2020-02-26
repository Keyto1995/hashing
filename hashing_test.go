package hashing

import (
	"bytes"
	"testing"
)

func TestMurmur32(t *testing.T) {
	buffer := bytes.NewBufferString("hello, world")
	hashFunction := Murmur32()
	hashCode := hashFunction.HashBytes(buffer.Bytes())

	t.Logf("编码长度：%v bits", hashCode.Bits())
	t.Logf("uint32: %v", hashCode.AsUint32())
	t.Logf("int32: %v", int32(hashCode.AsUint32()))
	t.Logf("bytes: %v", hashCode.AsBytes())
	t.Logf("hex: %v", hashCode.AsHex())
}

// TestMurmur32_1 结果与TestMurmur32()相同
func TestMurmur32_1(t *testing.T) {
	hasher := Murmur32().NewHasher()
	hasher.PutString("hello").PutByte(byte(',')).PutString(" world")
	hashCode := hasher.Hash()

	t.Logf("编码长度：%v bits", hashCode.Bits())
	t.Logf("uint32: %v", hashCode.AsUint32())
	t.Logf("int32: %v", int32(hashCode.AsUint32()))
	t.Logf("bytes: %v", hashCode.AsBytes())
	t.Logf("hex: %v", hashCode.AsHex())
}

func TestMurmur128(t *testing.T) {
	buffer := bytes.NewBufferString("hello, world")
	hashFunction := Murmur128()
	hashCode := hashFunction.HashBytes(buffer.Bytes())

	t.Logf("编码长度：%v bits", hashCode.Bits())
	t.Logf("uint32: %v", hashCode.AsUint32())
	t.Logf("int32: %v", int32(hashCode.AsUint32()))
	t.Logf("bytes: %v", hashCode.AsBytes())
	t.Logf("hex: %v", hashCode.AsHex())
}

// TestMurmur128_1 结果与TestMurmur128()相同
func TestMurmur128_1(t *testing.T) {
	hasher := Murmur128().NewHasher()
	hasher.PutString("hello").PutByte(byte(',')).PutString(" world")
	hashCode := hasher.Hash()

	t.Logf("编码长度：%v bits", hashCode.Bits())
	t.Logf("uint32: %v", hashCode.AsUint32())
	t.Logf("int32: %v", int32(hashCode.AsUint32()))
	t.Logf("bytes: %v", hashCode.AsBytes())
	t.Logf("hex: %v", hashCode.AsHex())
}
