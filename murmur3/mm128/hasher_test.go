package mm128

import (
	"reflect"
	"testing"
)

var hf = HashFunction{0}

func TestSum128WithSeed(t *testing.T) {
	type args struct {
		data []byte
		seed uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[]byte{1}, 0},
			want: "16fe7483905cce7a85670e43e4678877",
		},
		{
			name: "",
			args: args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 0},
			want: "95fc2048a0f04b08d038facbe83577a5",
		},
		{
			name: "",
			args: args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0},
			want: "e00f74ef6047c9b23589b912855d2f89",
		},
		{
			name: "",
			args: args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 0},
			want: "80acef949f88b31275fdfac9da346644",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hf := HashFunction{seed: tt.args.seed}
			if got := hf.NewHasher().PutBytes(tt.args.data[:1]).PutBytes(tt.args.data[1:]).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum128WithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "true",
			args: args{true},
			want: "16fe7483905cce7a85670e43e4678877",
		},
		{
			name: "false",
			args: args{false},
			want: "b55cff6ee5ab10468335f878aa2d6251",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutBool(tt.args.b).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutByte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{123},
			want: "e9a49f5b08d0e50558b8265bc51d3fee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutByte(tt.args.b).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[]byte{9, 11}},
			want: "0864028c8074f77edfec639834f94d6c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutBytes(tt.args.bytes).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutFloat32(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0.1",
			args: args{0.1},
			want: "66da88d68bea364277b7cf0d3bc6b937",
		},
		{
			name: "1.1",
			args: args{1.1},
			want: "fa2f42cbdd1264c5709a1a89773ff91b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutFloat32(tt.args.f).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutFloat64(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{1},
			want: "12b097ba1df0c93b0bf19e5d34594766",
		},
		{
			name: "1.1",
			args: args{1.1},
			want: "5293aad7197cc3edae606f47b410b060",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutFloat64(tt.args.f).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello,world",
			args: args{"hello,world"},
			want: "fa8830649f8a341239b7ee2641f3957b",
		},
		{
			name: "hello, world",
			args: args{"hello, world"},
			want: "8ebc5e3a62ac2f344d41429607bcdc4c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutString(tt.args.s).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutUint32(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{0},
			want: "bc764cd8ddf7a0cff126f51c16239658",
		},
		{
			name: "1",
			args: args{1},
			want: "feca28aff5a3958840bee985ee7de4d3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutUint32(tt.args.i).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_PutUint64(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{0},
			want: "cbc357ccb763df2852fee8c4fc7d55f2",
		},
		{
			name: "12",
			args: args{12},
			want: "1915c6ec29fddb53a9503e0321386482",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hf.NewHasher()
			if got := h.PutUint64(tt.args.i).Hash().AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
