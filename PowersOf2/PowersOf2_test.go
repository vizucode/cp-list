package powersof2

import (
	"reflect"
	"testing"
)

func TestPowersOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name     string
		args     args
		wantResp []uint
	}{
		{
			name: "Test 1",
			args: args{
				n: 0,
			},
			wantResp: []uint{1},
		},
		{
			name: "Test 2",
			args: args{
				n: 1,
			},
			wantResp: []uint{1, 2},
		},
		{
			name: "Test 3",
			args: args{
				n: 2,
			},
			wantResp: []uint{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := PowersOf2(tt.args.n); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PowersOf2() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
