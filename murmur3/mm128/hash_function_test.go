package mm128

import (
	"reflect"
	"testing"
)

func TestHashFunction_HashString(t *testing.T) {
	type fields struct {
		seed uint32
	}
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "",
			fields: fields{0},
			args:   args{"hello, world"},
			want:   "8ebc5e3a62ac2f344d41429607bcdc4c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hf := &HashFunction{
				seed: tt.fields.seed,
			}
			if got := hf.HashString(tt.args.data).AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashFunction_HashBytes(t *testing.T) {
	type fields struct {
		seed uint32
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{1}},
			want:   "16fe7483905cce7a85670e43e4678877",
		},

		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
			want:   "95fc2048a0f04b08d038facbe83577a5",
		},
		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
			want:   "e00f74ef6047c9b23589b912855d2f89",
		},
		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}},
			want:   "80acef949f88b31275fdfac9da346644",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hf := &HashFunction{
				seed: tt.fields.seed,
			}
			if got := hf.HashBytes(tt.args.data).AsHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
