package mm32

import (
	"github.com/keyto1995/hashing/hashcode"
	"reflect"
	"testing"
)

func TestMurmur3_32HashFunction_HashBytes(t *testing.T) {
	type fields struct {
		seed uint32
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   hashcode.HashCode
	}{
		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{}},
			want:   hashcode.NewUint32HashCode(0),
		},
		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{1, 2, 3, 4, 5, 6}},
			want:   hashcode.NewUint32HashCode(3619664446),
		},
		{
			name:   "",
			fields: fields{0},
			args:   args{[]byte{1, 2, 3, 4, 5, 7}},
			want:   hashcode.NewUint32HashCode(337367262),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &HashFunction{
				seed: tt.fields.seed,
			}
			if got := m.HashBytes(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
