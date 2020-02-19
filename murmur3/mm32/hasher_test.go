package mm32

import (
	"github.com/keyto1995/hashing/hashcode"
	"reflect"
	"testing"
)

func TestMurmur3_32Hasher_Hash(t *testing.T) {
	tests := []struct {
		name string
		want hashcode.HashCode
	}{
		{
			name: "",
			want: hashcode.NewUint32HashCode(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "PutBool(true)",
			args: args{true},
			want: hashcode.NewUint32HashCode(3831157163),
		},
		{
			name: "PutBool(false)",
			args: args{false},
			want: hashcode.NewUint32HashCode(1364076727),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutBool(tt.args.b).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutByte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "",
			args: args{1},
			want: hashcode.NewUint32HashCode(3831157163),
		},
		{
			name: "",
			args: args{255},
			want: hashcode.NewUint32HashCode(4251775245),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutByte(tt.args.b).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutBytes(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "",
			args: args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: hashcode.NewUint32HashCode(2711154856),
		},
		{
			name: "",
			args: args{[]byte{0, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: hashcode.NewUint32HashCode(3679365379),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutBytes(tt.args.bs).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutFloat32(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "",
			args: args{9.11},
			want: hashcode.NewUint32HashCode(1488918213),
		},
		{
			name: "",
			args: args{11.11},
			want: hashcode.NewUint32HashCode(4046293914),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutFloat32(tt.args.f).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutFloat64(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "",
			args: args{9.11},
			want: hashcode.NewUint32HashCode(2853424238),
		},
		{
			name: "",
			args: args{11.11},
			want: hashcode.NewUint32HashCode(3005471033),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutFloat64(tt.args.f).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "hello",
			args: args{"hello"},
			want: hashcode.NewUint32HashCode(613153351),
		},
		{
			name: "hel1o",
			args: args{"hel1o"},
			want: hashcode.NewUint32HashCode(2230949487),
		},
		{
			name: "你好",
			args: args{"你好"},
			want: hashcode.NewUint32HashCode(337357348),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutString(tt.args.s).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutUint32(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "",
			args: args{1234},
			want: hashcode.NewUint32HashCode(1807111040),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutUint32(tt.args.i).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMurmur3_32Hasher_PutUint64(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name string
		args args
		want hashcode.HashCode
	}{
		{
			name: "",
			args: args{996},
			want: hashcode.NewUint32HashCode(3463386159),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashFunction(0).NewHasher()
			if got := h.PutUint64(tt.args.i).Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
