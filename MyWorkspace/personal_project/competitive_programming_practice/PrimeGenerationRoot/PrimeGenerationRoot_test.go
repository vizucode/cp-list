package primegenerationroot

import (
	"reflect"
	"testing"
)

func TestPrimeGenerationRoot(t *testing.T) {
	type args struct {
		N uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "Test 1",
			args: args{
				N: 7,
			},
			want: []uint{2, 3, 5, 7},
		},
		{
			name: "Test 2",
			args: args{
				N: 11,
			},
			want: []uint{2, 3, 5, 7, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeGenerationRoot(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeGenerationRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
