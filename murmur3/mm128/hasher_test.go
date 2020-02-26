package mm128

import (
	"reflect"
	"testing"
)

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
