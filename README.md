# Hashing

提供流畅编程体验的哈希算法工具包

[ ok ] murmur3_32  
[  ] murmur3_128

- 安装包

```shell
go get github.com/keyto1995/hashing
```

- 导入包

``` go
import "github.com/keyto1995/hashing"
```

- `HashFunction` 是单纯的、无状态的方法，它把任意的数据块映射到固定数目的位指，并且保证相同的输入一定产生相同的输出，不同的输入尽可能产生不同的输出。 

``` go
hashFunction := hashing.Murmur32()
hashCode := hashFunction.HashString("hello, world")
```

- `HashFunction` 的实例可以提供有状态的 `Hasher`，`Hasher` 提供了流畅的语法把数据添加到散列运算，然后获取散列值。 

``` go
hasher := hashing.Murmur32().NewHasher()
hasher.PutString("hello").PutByte(byte(',')).PutString(" world")
hashCode := hasher.Hash()
```

- `HashCode` 保存了散列值，提供 `AsXxx()` 转化为对应的表现。

``` go
fmt.Printf("编码长度：%v bits\n", hashCode.Bits())
fmt.Printf("uint32: %v\n", hashCode.AsUint32())
fmt.Printf("int32: %v\n", int32(hashCode.AsUint32()))
fmt.Printf("bytes: %v\n", hashCode.AsBytes())
fmt.Printf("hex: %v\n", hashCode.AsHex())
```