package squarensum

import "testing"

func TestSquareNSum(t *testing.T) {
	type args struct {
		N []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				N: []int{1, 2, 2},
			},
			want: 9,
		},
		{
			name: "test case 2",
			args: args{
				N: []int{1, 2, 3},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareNSum(tt.args.N); got != tt.want {
				t.Errorf("SquareNSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
